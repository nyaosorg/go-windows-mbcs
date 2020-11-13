package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/zetamatta/go-windows-mbcs"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		text, err := mbcs.AtoU(sc.Bytes(), mbcs.ACP)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		fmt.Println(text)
	}
}
