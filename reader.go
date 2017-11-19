package store

type Reader interface {
	Finder
	Searcher
	Connector
	Pinger
	Closer
}
