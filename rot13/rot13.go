package main

import (
	"io"
	"os"
	"strings"
)

func isAlpha(char byte) bool {
	return (char > 96 && char < 123) || (char > 64 && char < 91)
}

func rot13(char byte) byte {
	var base byte

	if char > 96 {
		base = 95
	} else {
		base = 63
	}

	temp := char - base

	if temp > 13 {
		return base + temp%13
	} else {
		return char + 13
	}
}

type rot13Reader struct {
	r io.Reader
}

func (reader *rot13Reader) Read(buf []byte) (int, error) {
	var read_count int

	for {
		chunk_size, err := reader.r.Read(buf)
		read_count += chunk_size

		for idx := 0; idx < chunk_size; idx++ {
			if val := buf[idx]; isAlpha(val) {
				buf[idx] = rot13(val)
			}
		}

		if err != nil {
			return read_count, err
		}
	}
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := &rot13Reader{s}
	io.Copy(os.Stdout, r)
}
