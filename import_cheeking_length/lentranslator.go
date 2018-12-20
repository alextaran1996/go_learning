package main

import (
	"fmt"
	"os"
	"strconv"

	"./lenconv"
)

func main() {
	for _, i := range os.Args[1:] {
		t, err := strconv.ParseFloat(i, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		}
		arsh := lenconv.Metr(t)
		saz := lenconv.Metr(t)
		verst := lenconv.Metr(t)
		fmt.Printf("%s = %s, %s = %s, %s = %s\n", arsh, lenconv.MetrtoArsh(arsh), saz, lenconv.MetrtoSaz(saz), verst, lenconv.MetrtoVerst(verst))
	}
}
