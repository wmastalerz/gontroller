package controller

// Storage knows how to get the objects based on the ID, the Controller
// will get the object at the moment of handling the ID.
type Storage interface {
	Get(id string) (interface{}, error)
}

// StorageFunc is a helper to get a storer from a function
type StorageFunc func(id string) (interface{}, error)

// Get satisfies Storage interface.
func (s StorageFunc) Get(id string) (interface{}, error) {
	return s(id)
}
