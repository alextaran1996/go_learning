package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	var buf bytes.Buffer
	counter := len(s) % 3
	if counter == 0 {
		counter = 3
	}
	buf.WriteString(s[:counter])
	for i := counter; i < len(s); i += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[i : i+3])
	}
	return buf.String()
}

func main() {
	fmt.Println(comma("12345"))
}
