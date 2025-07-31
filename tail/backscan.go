package main

import (
	"bytes"
	"io"
)

// Inspired by icza/backscanner: https://github.com/icza/backscanner
// Basically a rip-off for learning purposes
// TODO: Can I make this work without a custom scanner?
type BackScan struct {
	r          io.ReaderAt // ReaderAt provided by client
	pos        int         // Offset, starting from EOF
	chunk_size int         // Read chunk_size bytes
	buf        []byte      // Read data
	err        error
}

func (s *BackScan) NewScanner(r io.ReaderAt, pos int) *BackScan {
	return &BackScan{r: r, pos: pos}
}

func (s *BackScan) readMore() {
	if s.pos == 0 {
		s.err = io.EOF
	}

	size := s.chunk_size
	// Don't read more than what's left to read
	if size > s.pos {
		size = s.pos
	}

	s.pos -= size

	// Read data into a new buf, then append existing data
	new_buf := make([]byte, size, size+len(s.buf))
	_, s.err = s.r.ReadAt(new_buf, int64(s.pos))

	if s.err == nil {
		s.buf = append(new_buf, s.buf...)
	}
}

func (s *BackScan) LineBytes() (line []byte, start int, err error) {
	if s.err != nil {
		return nil, 0, s.err
	}
	// First call on files with trailing newlines will return an empty line
	for {
		lineStart := bytes.LastIndexByte(s.buf, '\n')
		if lineStart > 0 {
			var line []byte
			line, s.buf = s.buf[lineStart+1:], s.buf[:lineStart]
			return line, s.pos + lineStart + 1, nil
		}
		s.readMore()
		if s.err != nil {
			if s.err == io.EOF && len(s.buf) > 0 {
				return s.buf, 0, nil
			}
			return nil, 0, s.err
		}
	}
}
