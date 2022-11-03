package pivot_server

import (
	"errors"
)

type Datastore interface {
	GetAllProductIDs() []int
	GetProductByID(int) (Product, error)
	GetInventoryByID(int) (Inventory, error)
	InsertProduct(Product) error
	UpdateProduct(Product) error
	InsertInventory(Inventory) error
	// ConditionallyUpdateInventory performs a compare-and-set operation, updating the inventory only if the current
	// value for the given item matches the supplied prev value. This is to prevent concurrent read-modify-write
	// operations from overwriting each other. Proper usage:
	// prev := ds.GetInventoryById(id)
	// updated := prev
	// updated.Inventory = prev.Inventory + 10
	// ds.ConditionallyUpdateInventory(prev, updated)
	ConditionallyUpdateInventory(prev Inventory, new Inventory) error
}

var ErrNotFound = errors.New("not found")
var ErrUpdateFailed = errors.New("udpate failed")
var ErrInsertFailed = errors.New("insert failed")
