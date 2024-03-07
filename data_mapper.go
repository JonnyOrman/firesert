package firesert

type DataMapper[TFrom any, TTo any] interface {
	Map(dataFrom TFrom) TTo
}
