package store

type Updater interface {
	Update(id int, obj map[string]interface{}) error
}
