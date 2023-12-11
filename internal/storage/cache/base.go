package cache

type Base interface {
	Get(string) bool
	Set(string)
	Delete(string) bool
}
