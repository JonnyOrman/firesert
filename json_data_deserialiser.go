package firesert

import "encoding/json"

type JsonDataDeserialiser[T any] struct{}

func (this JsonDataDeserialiser[T]) Deserialise(data []byte) T {
	var t T
	json.Unmarshal([]byte(string(data)), &t)

	return t
}
