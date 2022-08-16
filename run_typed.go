package firesert

func RunTyped[T any]() {
	application := BuildApplication[T]()

	application.Run()
}
