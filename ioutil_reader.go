package firesert

import (
	"io"
	"io/ioutil"
)

type IoutilReader struct{}

func (this IoutilReader) Read(ioReader io.Reader) []byte {
	data, _ := ioutil.ReadAll(ioReader)

	return data
}
