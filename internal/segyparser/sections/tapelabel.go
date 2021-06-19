package sections

import (
	"bufio"
	"errors"
)

// Type label is a 128 byte section at the front of segy file (optional)

const typeLabelSize = 128

type TapeLabel struct {
	data []byte
}

func (tl *TapeLabel) Data() []byte {
	res := make([]byte, typeLabelSize)
	copy(res, tl.data)
	return res
}

func (tl *TapeLabel) Read(r *bufio.Reader) error {
	tl.data = make([]byte, typeLabelSize)
	n, err := r.Read(tl.data)
	if err != nil {
		return err
	}
	if n != typeLabelSize {
		return errors.New("can't read Tape Label since the number of bytes is less then expected")
	}

	return nil
}
