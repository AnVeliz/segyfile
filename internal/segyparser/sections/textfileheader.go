package sections

// Textual File Header is a 3200 byte section after tape label

type TextualFileHeader struct {
	data [3200]byte
}

func (tl TextualFileHeader) Data() [3200]byte {
	return tl.data
}
