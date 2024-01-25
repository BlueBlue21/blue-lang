package cli

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func setup() {
	flag.Usage = func() {
		fmt.Printf("Usage: %s [ code.bl ]\n\n", os.Args[0])
	}
	flag.Parse()
}

func checkFile(file string) {
	_, err := os.Stat(file)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File not found: " + file)
	} else {
		dat, _ := ioutil.ReadFile(file)
		fmt.Println(string(dat))
	}
}

func Run() {
	setup()
	if len(flag.Args()) == 1 {
		checkFile(flag.Arg(0))
	} else {
		flag.Usage()
	}
}
