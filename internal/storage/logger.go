package storage

type Logger interface {
	Printf(string, ...interface{})
}
