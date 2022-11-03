package pivot_server

import (
	"net/http"
)

type MyInventoryServer struct {
	ds Datastore
}

// HandleGetInventories should handle GET /inventories requests
func (m *MyInventoryServer) HandleGetInventories(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	w.WriteHeader(500)
}

// HandleGetInventory should handle GET /inventory/{product_id} requests
func (m *MyInventoryServer) HandleGetInventory(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	w.WriteHeader(500)
}

// HandleCreateInventory should handle PUT /inventories/{product_id} requests
func (m *MyInventoryServer) HandleCreateInventory(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	w.WriteHeader(500)
}

// HandleUpdateInventory should handle POST /inventories/{product_id} requests
func (m *MyInventoryServer) HandleUpdateInventory(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	w.WriteHeader(500)
}
