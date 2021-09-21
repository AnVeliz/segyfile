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

func (d *RawData) Data() []byte {
	dataSize := binary.Size(d.data)
	if dataSize == 0 || dataSize == -1 {
		return nil
	}

	result := make([]byte, dataSize)
	copy(result, d.data)
	return d.data
}

func (d *RawData) Read(reader *bufio.Reader, resultObject interface{}) error {
	objectSize := binary.Size(resultObject)
	d.data = make([]byte, objectSize)
	readLen, err := reader.Read(d.data)

	if err != nil {
		return err
	}
	if readLen < objectSize {
		return errors.New("can't read data chunk since the number of bytes read is less then expected")
	}

	err = binary.Read(bytes.NewReader(d.data), binary.BigEndian, resultObject)
	if err != nil {
		return err
	}

	return nil
}
