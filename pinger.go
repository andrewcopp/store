package store

type Pinger interface {
	Ping() error
}
