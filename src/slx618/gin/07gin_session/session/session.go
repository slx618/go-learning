package session

type Session interface {
	Set(key string, val interface{}) (err error)

	Get(key string) (interface{}, error)

	Del(key string) error

	Save() error
}
