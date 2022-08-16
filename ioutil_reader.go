package firesert

import (
	"io"
	"io/ioutil"
)

type IoutilReader struct{}

func (ioutilReader IoutilReader) Read(ioReader io.Reader) []byte {
	data, _ := ioutil.ReadAll(ioReader)

	return data
}
