package sections

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

// Binary File Header is a 400 byte section after textual file header

const binaryFileHeaderSize = 400

type BinaryFileHeader struct {
	JobId int32
	Lino  int32
	Reno  int32
	Ntrpr int16
}

type BinaryFileHeaderRaw struct {
	data []byte
}

func (bh *BinaryFileHeaderRaw) Data() []byte {
	res := make([]byte, binaryFileHeaderSize)
	copy(res, bh.data)
	return bh.data
}

func (bh *BinaryFileHeaderRaw) Read(r *bufio.Reader) error {
	bh.data = make([]byte, binaryFileHeaderSize)
	n, err := r.Read(bh.data)
	if err != nil {
		return err
	}
	if n != binaryFileHeaderSize {
		return errors.New("can't read Tape Label since the number of bytes is less then expected")
	}

	var hf BinaryFileHeader
	err = binary.Read(bytes.NewReader(bh.data), binary.BigEndian, &hf)
	if err != nil {
		return err
	}

	fmt.Printf("Res: %v\n", hf)

	return nil
}
