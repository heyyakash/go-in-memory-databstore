package datastore

import "sync"

type DataStore struct {
	mutex sync.RWMutex
	Store map[string]string
}

var KeyValueStore *DataStore

func CreateNewStore() {
	KeyValueStore = &DataStore{
		Store: make(map[string]string),
	}
}

func (d *DataStore) SetValue(key string, value string) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	d.Store[key] = value
}

func (d *DataStore) GetValue(key string) (string, bool) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	val, ok := d.Store[key]
	return val, ok
}

func (d *DataStore) DeleteValue(key string) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	delete(d.Store, key)
}

func (d *DataStore) KeyExists(key string) bool {
	_, exists := d.Store[key]
	return exists
}
