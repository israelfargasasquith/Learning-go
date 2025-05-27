package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	if err != nil {
		return -1, fmt.Errorf("Error with reader")
	}

	for i := 0; i < n; i++ {
		switch {
		case b[i] >= 'A' && b[i] <= 'Z':
			b[i] = 'A' + (b[i]-'A'+13)%26
		case b[i] >= 'a' && b[i] <= 'z':
			b[i] = 'a' + (b[i]-'a'+13)%26
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
