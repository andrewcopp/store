package store

type Finder interface {
	Find(id int) (map[string]interface{}, error)
}
