package pivot_server

import (
	"encoding/csv"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// ProductInventory is a useful struct for some requests and responses
type ProductInventory struct {
	ProductID int       `json:"product_id"`
	Name      string    `json:"name"`
	Inventory int       `json:"inventory"`
	UpdatedAt time.Time `json:"updated_at"`
}

type InventoryAdjustment struct {
	Adjustment int `json:"inventory_adjustment"`
}

type InventoryServer interface {
	HandleGetInventories(w http.ResponseWriter, r *http.Request)
	HandleGetInventory(w http.ResponseWriter, r *http.Request)
	HandleCreateInventory(w http.ResponseWriter, r *http.Request)
	HandleUpdateInventory(w http.ResponseWriter, r *http.Request)
}

func InventoryRouter(s InventoryServer) http.HandlerFunc {
	allInventories := regexp.MustCompile("^/inventories(/)?$")
	itemInventory := regexp.MustCompile("^/inventories/[\\w\\d]+")
	routes := []struct {
		method  string
		re      *regexp.Regexp
		handler func(w http.ResponseWriter, r *http.Request)
	}{
		{
			method:  "GET",
			re:      allInventories,
			handler: s.HandleGetInventories,
		},
		{
			method:  "PUT",
			re:      itemInventory,
			handler: s.HandleCreateInventory,
		},
		{
			method:  "GET",
			re:      itemInventory,
			handler: s.HandleGetInventory,
		},
		{
			method:  "POST",
			re:      itemInventory,
			handler: s.HandleUpdateInventory,
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		for _, route := range routes {
			if route.method == r.Method && route.re.MatchString(r.URL.Path) {
				route.handler(w, r)
				return
			}
		}
		w.WriteHeader(404)
	}
}

func LoadSeedData(ds Datastore) error {
	return LoadCSV(ds, `1,R2D2,100,2022-01-01T00:00:00Z
2,Millennium Falcon,200,2022-01-02T00:00:00Z
3,Darth Vader,50,2022-01-03T00:00:00Z
4,Yoda,10,2022-01-04T00:00:00Z`)
}

func LoadCSV(ds Datastore, data string) error {
	r := csv.NewReader(strings.NewReader(data))
	r.FieldsPerRecord = 4
	rs, err := r.ReadAll()
	if err != nil {
		return err
	}
	// fields: id, name, inventory, updated_at
	for _, rec := range rs {
		id, err := strconv.Atoi(rec[0])
		if err != nil {
			return err
		}
		name := rec[1]
		inventory, err := strconv.Atoi(rec[2])
		if err != nil {
			return err
		}
		dt, err := time.Parse(time.RFC3339, rec[3])
		if err != nil {
			return err
		}
		p := Product{
			ID:   id,
			Name: name,
		}
		i := Inventory{
			ID:        id,
			Inventory: inventory,
			UpdatedAt: dt,
		}
		if err := ds.InsertProduct(p); err != nil {
			return err
		}
		if err := ds.InsertInventory(i); err != nil {
			return err
		}
	}
	return nil
}
