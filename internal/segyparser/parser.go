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

func Parse(fileName string, parseScope ParseScope) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error loading the file")
		return
	}

	bufReader := bufio.NewReader(file)
	data := make([]byte, sections.TEXTUAL_HEADER_LEN)
	bufReader.Read(data)
	binaryHeader := sections.BinaryFileHeader{}
	binaryHeader.Read(bufReader)
}
