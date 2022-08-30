package firesert

type ValueProvider interface {
	Get() (string, bool)
}
