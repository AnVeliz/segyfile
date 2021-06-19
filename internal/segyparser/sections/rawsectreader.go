package sections

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
)

type RawData struct {
	data []byte
}

func (bh *RawData) Data() []byte {
	res := make([]byte, binaryFileHeaderSize)
	copy(res, bh.data)
	return bh.data
}

func (bh *RawData) Read(r *bufio.Reader, resObj interface{}) error {
	objSize := binary.Size(resObj)
	bh.data = make([]byte, objSize)
	n, err := r.Read(bh.data)

	if err != nil {
		return err
	}
	if n < objSize {
		return errors.New("can't read data chunk since the number of bytes read is less then expected")
	}

	err = binary.Read(bytes.NewReader(bh.data), binary.BigEndian, resObj)
	if err != nil {
		return err
	}

	return nil
}
