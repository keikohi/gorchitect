package registry

import (
	"io"

	"github.com/keikohi/gorchitect/domain/repository"
	"github.com/keikohi/gorchitect/persistence/writer"
)

type Writer interface {
	NewWriter() repository.Writer
}

type WriterImpl struct {
	writer io.Writer
}

func NewWriter(writer io.Writer) WriterImpl {
	return WriterImpl{writer: writer}
}

func (w WriterImpl) NewWriter() repository.Writer {
	return writer.Dotwriter{W: w.writer}
}
