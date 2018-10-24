package datastore

type DataStore interface {
	Open() error
	Init() error
	Close() error
}
