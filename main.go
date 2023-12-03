package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
)

func main() {
	h := md5.New()
	io.WriteString(h, "This class rocks")
	fmt.Printf("%x", h.Sum(nil))

	h = sha1.New()
	io.WriteString(h, "This class rocks")
	fmt.Printf("% x", h.Sum(nil))
}
