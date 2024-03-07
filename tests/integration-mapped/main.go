package main

import "github.com/jonnyorman/firesert"

type MessageData struct {
	Prop1 string
	Prop2 int
	Prop3 string
}

type TestDocument struct {
	Prop1 string
	Prop2 int
}

type MessageDataToTestDocumentMapper struct{}

func (this MessageDataToTestDocumentMapper) Map(from MessageData) TestDocument {
	return TestDocument{
		Prop1: from.Prop1,
		Prop2: from.Prop2,
	}
}

func main() {
	mapper := MessageDataToTestDocumentMapper{}

	firesert.RunMapped[MessageData, TestDocument](mapper)
}
