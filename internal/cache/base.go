package cache

type ICache interface {
	Get(string) bool
	Set(string)
	Delete(string) bool
}
