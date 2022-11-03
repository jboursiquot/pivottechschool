package pivot_server

import (
	"reflect"
	"sync"
	"time"
)

func (m *MemoryInventoryStore) GetAllProductIDs() []int {
	m.mu.Lock()
	defer m.mu.Unlock()
	var out []int
	for k := range m.products {
		out = append(out, k)
	}
	return out
}

func (m *MemoryInventoryStore) GetProductByID(id int) (Product, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	p, ok := m.products[id]
	if !ok {
		return Product{}, ErrNotFound
	}
	return p, nil
}

func (m *MemoryInventoryStore) GetInventoryByID(id int) (Inventory, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	p, ok := m.inventories[id]
	if !ok {
		return Inventory{}, ErrNotFound
	}
	return p, nil
}

func (m *MemoryInventoryStore) InsertProduct(product Product) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	_, ok := m.products[product.ID]
	if ok {
		return ErrInsertFailed
	}
	m.products[product.ID] = product
	return nil
}

func (m *MemoryInventoryStore) UpdateProduct(product Product) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	_, ok := m.products[product.ID]
	if !ok {
		return ErrUpdateFailed
	}
	m.products[product.ID] = product
	return nil
}

func (m *MemoryInventoryStore) InsertInventory(inventory Inventory) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	_, ok := m.inventories[inventory.ID]
	if ok {
		return ErrInsertFailed
	}
	m.inventories[inventory.ID] = inventory
	return nil
}

func (m *MemoryInventoryStore) ConditionallyUpdateInventory(prev Inventory, new Inventory) error {
	if prev.ID != new.ID {
		return ErrUpdateFailed
	}

	m.mu.Lock()
	defer m.mu.Unlock()
	actualPrev, ok := m.inventories[prev.ID]
	if !ok {
		return ErrUpdateFailed
	}
	if !reflect.DeepEqual(prev, actualPrev) {
		return ErrUpdateFailed
	}
	m.inventories[new.ID] = new
	return nil
}

type Product struct {
	ID   int
	Name string
}

type Inventory struct {
	ID        int
	Inventory int
	UpdatedAt time.Time
}

type MemoryInventoryStore struct {
	products    map[int]Product
	inventories map[int]Inventory
	mu          sync.Mutex
}

func NewMemoryDatastore() *MemoryInventoryStore {
	return &MemoryInventoryStore{
		products:    make(map[int]Product),
		inventories: make(map[int]Inventory),
		mu:          sync.Mutex{},
	}
}
