package sections

// Textual File Header is a 3200 byte section after tape label
const TEXTUAL_HEADER_LEN = 3200

type TextualFileHeader struct {
	data [TEXTUAL_HEADER_LEN]byte
}

func (h TextualFileHeader) Data() [3200]byte {
	return h.data
}
