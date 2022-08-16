package firesert

type DataInserter[T any] interface {
	Insert(data T)
}
