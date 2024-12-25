package database

type IDB interface {
	Open() error
	Close()
}
