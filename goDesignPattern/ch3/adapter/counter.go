package main

import (
	"io"
	"log"
	"os"
	"strconv"
)

type counter struct {
	io.Writer
}

func (c *counter) count(n uint64) uint64 {
	if n == 0 {
		c.Write([]byte(strconv.Itoa(0) + "\n"))
		return 0
	}
	cur := n
	c.Write([]byte(strconv.FormatUint(cur, 10) + "\n"))
	return c.count(n - 1)
}
func main() {
	pipeReader, pipeWriter := io.Pipe()
	defer pipeWriter.Close()
	defer pipeReader.Close()

	c := counter{pipeWriter}

	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	tee := io.TeeReader(pipeReader, f)

	go func() {
		io.Copy(os.Stdout, tee)
	}()

	c.count(8)

}
