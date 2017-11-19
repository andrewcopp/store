package store

type Deleter interface {
	Delete(id int) error
}
