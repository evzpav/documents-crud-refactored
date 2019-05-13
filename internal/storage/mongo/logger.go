package mongo

type logger interface {
	Printf(string, ...interface{})
}
