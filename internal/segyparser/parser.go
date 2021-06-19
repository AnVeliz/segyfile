package segyparser

import (
	"bufio"
	"fmt"
	"os"

	"github.com/AnVeliz/segyfile/internal/segyparser/sections"
)

type ParseScope int

const (
	ParseEverything ParseScope = iota
)

func Parse(file string, scope ParseScope) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Error loading the file")
		return
	}

	r := bufio.NewReader(f)
	d := make([]byte, 3200)
	r.Read(d)
	hr := sections.BinaryFileHeader{}
	hr.Read(r)
}
