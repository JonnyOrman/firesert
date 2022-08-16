package firesert

import "encoding/json"

type JsonDataDeserialiser[T any] struct {
}

func (jsonDataDeserialiser JsonDataDeserialiser[T]) Deserialise(data []byte) T {
	var t T
	json.Unmarshal([]byte(string(data)), &t)

	return t
}
