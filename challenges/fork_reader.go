package challenges

import (
	"io"
)

func ForkReader(r io.Reader, w1, w2 io.Writer) error {
	buffer := make([]byte, 1)
	writers := []io.Writer{w1, w2}
	turn := 0

	for {
		n, err := r.Read(buffer)
		if n > 0 {
			if _, wErr := writers[turn%2].Write(buffer); wErr != nil {
				return wErr
			}
			turn++
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
	}
	return nil
}
