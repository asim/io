package pipe

import (
	"bufio"
	"io"
)

type Pipe struct {
	reader io.Reader
	writer io.Writer
	buffer *bufio.Writer
}

func (p *Pipe) Close() error {
	if err := p.reader.(*io.PipeReader).Close(); err != nil {
		return err
	}
	if err := p.writer.(*io.PipeWriter).Close(); err != nil {
		return err
	}
	return nil
}

func (p *Pipe) Read(data []byte) (int, error) {
	return p.reader.Read(data)
}

func (p *Pipe) Write(data []byte) (int, error) {
	return p.buffer.Write(data)
}

func New() (*Pipe, error) {
	r, w := io.Pipe()

	return &Pipe{
		reader: r,
		writer: w,
		buffer: bufio.NewWriter(w),
	}, nil
}
