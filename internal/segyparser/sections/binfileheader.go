package sections

import (
	"bufio"
)

// Binary File Header is a 400 byte section after textual file header

const binaryFileHeaderSize = 400

type BinaryFileHeaderContent struct {
	JobId int32
	Lino  int32
	Reno  int32
	Ntrpr int16
}
type BinaryFileHeader struct {
	rawContent RawData
	content    BinaryFileHeaderContent
}

func (bh *BinaryFileHeader) Read(r *bufio.Reader) error {
	bh.rawContent = RawData{}
	bh.rawContent.Read(r, &bh.content)

	return nil
}
