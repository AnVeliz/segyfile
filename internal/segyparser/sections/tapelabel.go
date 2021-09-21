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

func (l *TapeLabel) Data() []byte {
	res := make([]byte, typeLabelSize)
	copy(res, l.data)
	return res
}

func (l *TapeLabel) Read(r *bufio.Reader) error {
	l.data = make([]byte, typeLabelSize)
	n, err := r.Read(l.data)
	if err != nil {
		return err
	}
	if n != typeLabelSize {
		return errors.New("can't read Tape Label since the number of bytes is less then expected")
	}

	return nil
}
