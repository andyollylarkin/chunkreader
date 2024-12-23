package chunkreader

import (
	"bytes"
	"io"
)

type ChunkReader struct {
	from      io.Reader
	chunkSize int
	buffer    *bytes.Buffer
	// buffer for reducing system calls
	currentBuffRead int64
}

func NewChunkReader(readFrom io.Reader, chunkSize int) *ChunkReader {
	return &ChunkReader{
		from:      readFrom,
		chunkSize: chunkSize,
		buffer:    bytes.NewBuffer(make([]byte, 0)),
	}
}

func (r *ChunkReader) Read(p []byte) (int, error) {
	var err error

	if r.currentBuffRead <= 0 {
		r.currentBuffRead, err = io.Copy(r.buffer, io.LimitReader(r.from, int64(r.chunkSize)))
		if err != nil {
			return int(r.currentBuffRead), err
		}
	}

	readed, err := r.buffer.Read(p)
	if err != nil {
		return readed, err
	}

	r.currentBuffRead -= int64(readed)

	if r.currentBuffRead <= 0 {
		r.currentBuffRead = 0
		r.buffer.Reset()
	}

	return readed, err
}
