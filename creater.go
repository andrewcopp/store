package store

type Creater interface {
	Create(obj map[string]interface{}) error
}
