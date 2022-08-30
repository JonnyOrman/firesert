package firesert

type JsonReader interface {
	Read() map[string]interface{}
}
