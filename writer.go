package store

type Writer interface {
	Reader
	Creater
	Updater
	Deleter
}
