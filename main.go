package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	fmt.Printf("Message Digest Version 5 MD5\n")
	h := md5.New()
	plaintext := "This class rocks"
	fmt.Printf(plaintext + "\n")
	io.WriteString(h, plaintext)
	fmt.Printf("%x\n", h.Sum(nil))
	fmt.Printf("The number of bytes: %d\n", h.Size())

	fmt.Printf("\nSecure Hash Algorithm Version 1 SHA\n")
	fmt.Printf(plaintext + "\n")
	h = sha1.New()
	io.WriteString(h, plaintext)
	fmt.Printf("%x\n", h.Sum(nil))
	fmt.Printf("The number of bytes: %d\n", h.Size())

	fmt.Printf("\nSecure Hash Algorithm Version 2 SHA256\n")
	s := "{'payload':'hashLab'}"
	h = sha256.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
	fmt.Printf("The number of bytes: %d\n", h.Size())
}
