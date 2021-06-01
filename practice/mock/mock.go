package mock

type Mock interface {
	Get(key string) string
}
