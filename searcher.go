package store

type Searcher interface {
	Search(off int, lim int) ([]map[string]interface{}, error)
}
