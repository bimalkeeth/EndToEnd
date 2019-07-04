package set

type Interface interface {
	Add(v interface{}) bool
	Remove(v interface{}) bool
	IsElementOf(v interface{}) bool
	Size() int
}
type Emptier interface {
	Empty()
}
