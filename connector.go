package store

type Connector interface {
	Connect() error
}
