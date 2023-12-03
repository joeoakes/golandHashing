package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
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

	s := "{'payload':'hashLab'}"
	h = sha256.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
