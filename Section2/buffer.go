package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

func main() {
	_ = WorkWithBuffer()
}

func Buffer(rawString string) *bytes.Buffer {
	rawBytes := []byte(rawString)
	var b = new(bytes.Buffer)
	b.Write(rawBytes)
	b = bytes.NewBuffer(rawBytes)
	b = bytes.NewBufferString(rawString)
	return b

}

func toString(r io.Reader) (string, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func WorkWithBuffer() error {
	rawString := "it is easy to encode unicode into a byte array"

	b := Buffer(rawString)
	fmt.Println(b)
	s, err := toString(b)
	if err != nil {
		return err
	}
	fmt.Println(s)
	reader := bytes.NewReader([]byte(rawString))
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return nil
}
