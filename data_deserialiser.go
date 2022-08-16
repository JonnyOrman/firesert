package firesert

type DataDeserialiser[T any] interface {
	Deserialise(data []byte) T
}
