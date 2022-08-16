package firesert

import (
	"bytes"
	"io"
)

type ConfigurationIoReaderGenerator struct{}

func (this ConfigurationIoReaderGenerator) Generate(payload string) io.Reader {
	return bytes.NewBufferString(payload)
}
