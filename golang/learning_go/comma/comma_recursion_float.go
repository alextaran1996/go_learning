package main

import "os"

func commarecursfl(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return commarecursfl(s[:n-3]) + "," + s[n-3:]

}

func procosarg() {
	for _, val := range os.Args[1:] {
		result := commarecursfl(val)
		println(result)
	}
}

func main() {
	procosarg()
}
