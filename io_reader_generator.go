package firesert

import "io"

type IoReaderGenerator interface {
	Generate(payload string) io.Reader
}
