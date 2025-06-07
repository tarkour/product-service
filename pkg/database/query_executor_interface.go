package db

type Executor interface {
	Execute(query string) (string, error)
}
