package utilities

import (
	"io"
)

// Rot128Reader implements io.Reader that transforms
type Rot128Reader struct {
	reader io.Reader
}

// NewRot128Reader for initial reader
func NewRot128Reader(r io.Reader) (*Rot128Reader, error) {
	return &Rot128Reader{reader: r}, nil
}

// Read for decrypt/reverse
func (r *Rot128Reader) Read(p []byte) (int, error) {
	n, err := r.reader.Read(p)
	if err != nil {
		return n, err
	}

	rot128(p[:n])
	return n, nil
}

// Rot128Writer implements io.Writer that transforms
type Rot128Writer struct {
	writer io.Writer
	buffer []byte // not thread-safe
}

// NewRot128Writer for initial writer
func NewRot128Writer(w io.Writer) (*Rot128Writer, error) {
	return &Rot128Writer{
		writer: w,
		buffer: make([]byte, 4096),
	}, nil
}

// Write for decrypt/reverse
func (w *Rot128Writer) Write(p []byte) (int, error) {
	n := copy(w.buffer, p)
	rot128(w.buffer[:n])
	return w.writer.Write(w.buffer[:n])
}

func rot128(buf []byte) {
	for idx := range buf {
		buf[idx] += 128
	}
}
