package firesert

import (
	"bytes"
	"io"
)

type ConfigurationIoReaderGenerator struct{}

func (configurationIoReaderGenerator ConfigurationIoReaderGenerator) Generate(payload string) io.Reader {
	return bytes.NewBufferString(payload)
}
