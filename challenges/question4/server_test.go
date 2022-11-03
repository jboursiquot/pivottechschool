package pivot_server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func createMyServer() http.Handler {
	ds := NewMemoryDatastore()
	if err := LoadSeedData(ds); err != nil {
		panic(err)
	}
	s := MyInventoryServer{ds}
	return InventoryRouter(&s)
}

func TestGetAll(t *testing.T) {
	server := createMyServer()
	t.Run("ok", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/inventories", nil)
		server.ServeHTTP(rr, req)
		assertStatusCode(t, rr, 200)
		var out []ProductInventory
		err := json.NewDecoder(rr.Result().Body).Decode(&out)
		assertNoError(t, err)
		assertEquals(t, 4, len(out))
	})
}

func TestGetOne(t *testing.T) {
	server := createMyServer()
	t.Run("ok", func(t *testing.T) {
		out := getOne(t, server, 3)
		assertEquals(t, 3, out.ProductID)
		assertEquals(t, "Darth Vader", out.Name)
		assertEquals(t, 50, out.Inventory)
	})
	t.Run("error not found", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/inventories/7", nil)
		server.ServeHTTP(rr, req)
		assertStatusCode(t, rr, 404)
	})
}

func TestCreateInventory(t *testing.T) {
	server := createMyServer()
	t.Run("ok", func(t *testing.T) {
		rr := httptest.NewRecorder()
		pi := ProductInventory{
			ProductID: 7,
			Name:      "Mando",
			Inventory: 71,
			UpdatedAt: time.Now().UTC(),
		}
		body, _ := json.Marshal(pi)
		req, _ := http.NewRequest("PUT", "/inventories/7", bytes.NewBuffer(body))
		server.ServeHTTP(rr, req)
		assertStatusCode(t, rr, 201)
		var out ProductInventory
		err := json.NewDecoder(rr.Result().Body).Decode(&out)
		assertNoError(t, err)
		assertEquals(t, pi, out)
	})
	t.Run("error bad request - invalid JSON", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", "/inventories/8", bytes.NewBuffer([]byte("{")))
		server.ServeHTTP(rr, req)
		assertStatusCode(t, rr, 400)
	})
	t.Run("error bad request - duplicate item", func(t *testing.T) {
		rr := httptest.NewRecorder()
		pi := ProductInventory{
			ProductID: 3,
			Name:      "Mando",
			Inventory: 71,
			UpdatedAt: time.Now().UTC(),
		}
		body, _ := json.Marshal(pi)
		req, _ := http.NewRequest("PUT", "/inventories/3", bytes.NewBuffer(body))
		server.ServeHTTP(rr, req)
		assertStatusCode(t, rr, 409)
	})
}

func TestUpdateInventory(t *testing.T) {
	server := createMyServer()
	t.Run("ok - increment", func(t *testing.T) {
		rr := httptest.NewRecorder()
		adj := InventoryAdjustment{2}
		body, _ := json.Marshal(adj)
		req, _ := http.NewRequest("POST", "/inventories/1", bytes.NewBuffer(body))
		server.ServeHTTP(rr, req)
		assertStatusCode(t, rr, 200)
		var out ProductInventory
		err := json.NewDecoder(rr.Result().Body).Decode(&out)
		assertNoError(t, err)
		assertEquals(t, 102, out.Inventory)
	})
	t.Run("ok - decrement", func(t *testing.T) {
		rr := httptest.NewRecorder()
		adj := InventoryAdjustment{-2}
		body, _ := json.Marshal(adj)
		req, _ := http.NewRequest("POST", "/inventories/2", bytes.NewBuffer(body))
		server.ServeHTTP(rr, req)
		assertStatusCode(t, rr, 200)
		var out ProductInventory
		err := json.NewDecoder(rr.Result().Body).Decode(&out)
		assertNoError(t, err)
		assertEquals(t, 198, out.Inventory)
	})
	t.Run("error bad request - invalid JSON", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/inventories/2", bytes.NewBuffer([]byte("{")))
		server.ServeHTTP(rr, req)
		assertStatusCode(t, rr, 400)
	})
	t.Run("error bad request - invalid value", func(t *testing.T) {
		rr := httptest.NewRecorder()
		adj := InventoryAdjustment{-2000}
		body, _ := json.Marshal(adj)
		req, _ := http.NewRequest("POST", "/inventories/3", bytes.NewBuffer(body))
		server.ServeHTTP(rr, req)
		assertStatusCode(t, rr, 400)
		after := getOne(t, server, 3)
		assertEquals(t, 50, after.Inventory)
	})
}

func getOne(t *testing.T, server http.Handler, id int) ProductInventory {
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/inventories/%d", id), nil)
	server.ServeHTTP(rr, req)
	assertStatusCode(t, rr, 200)
	var out ProductInventory
	err := json.NewDecoder(rr.Result().Body).Decode(&out)
	assertNoError(t, err)
	return out
}
