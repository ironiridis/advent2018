package justscan

import (
	"bufio"
	"os"
)

// Chan returns a channel where sequential lines of a file
// are returned. It will close the channel when all lines
// are read. It panics on any errors.
func Chan(fn string) (c chan string) {
	fp, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(fp)
	c = make(chan string)
	go func() {
		for s.Scan() {
			c <- s.Text()
		}
		if err := s.Err(); err != nil {
			panic(err)
		}
		fp.Close()
		close(c)
	}()
	return c
}
