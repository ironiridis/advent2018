package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func must(s string, e error) {
	if e != nil {
		fmt.Printf("failed to %s: %s\n", s, e)
		panic(e)
	}
}

var pair = map[byte]byte{
	0x41: 0x61, 0x61: 0x41,
	0x42: 0x62, 0x62: 0x42,
	0x43: 0x63, 0x63: 0x43,
	0x44: 0x64, 0x64: 0x44,
	0x45: 0x65, 0x65: 0x45,
	0x46: 0x66, 0x66: 0x46,
	0x47: 0x67, 0x67: 0x47,
	0x48: 0x68, 0x68: 0x48,
	0x49: 0x69, 0x69: 0x49,
	0x4a: 0x6a, 0x6a: 0x4a,
	0x4b: 0x6b, 0x6b: 0x4b,
	0x4c: 0x6c, 0x6c: 0x4c,
	0x4d: 0x6d, 0x6d: 0x4d,
	0x4e: 0x6e, 0x6e: 0x4e,
	0x4f: 0x6f, 0x6f: 0x4f,
	0x50: 0x70, 0x70: 0x50,
	0x51: 0x71, 0x71: 0x51,
	0x52: 0x72, 0x72: 0x52,
	0x53: 0x73, 0x73: 0x53,
	0x54: 0x74, 0x74: 0x54,
	0x55: 0x75, 0x75: 0x55,
	0x56: 0x76, 0x76: 0x56,
	0x57: 0x77, 0x77: 0x57,
	0x58: 0x78, 0x78: 0x58,
	0x59: 0x79, 0x79: 0x59,
	0x5a: 0x7a, 0x7a: 0x5a,
}

func reactRead(polyin io.Reader) []byte {
	buf := make([]byte, 1)
	polyout := new(bytes.Buffer)

	for {
		_, err := polyin.Read(buf)
		if err == io.EOF {
			break // done!
		}
		must("read polymer byte", err)

		if polyout.Len() == 0 {
			polyout.WriteByte(buf[0])
			continue
		}

		if pair[polyout.Bytes()[polyout.Len()-1]] == buf[0] {
			polyout.Truncate(polyout.Len() - 1) // back up one position
		} else {
			polyout.WriteByte(buf[0])
		}
	}
	return polyout.Bytes()
}

func main() {
	f, err := os.Open("../input.txt")
	must("open input file", err)
	k := reactRead(f)
	f.Close()
	fmt.Printf("k: %q (%d bytes)\n", k, len(k))
}
