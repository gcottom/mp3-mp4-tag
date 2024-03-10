package audiometa

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"strings"
)

// GetFileType returns the file type of the file pointed to by filepath. If the filetype is not supported, an error is returned.
func GetFileType(filepath string) (FileType, error) {
	sp := strings.Split(filepath, ".")
	if len(sp) < 2 {
		return "", errors.New("unsupported file extension or no extension")
	}
	for _, ft := range supportedFileTypes {
		if strings.ToLower(sp[len(sp)-1]) == string(ft) {
			return ft, nil
		}
	}
	return "", errors.New("unsupported file extension or no extension")
}

func getInt(b []byte) int {
	var n int
	for _, x := range b {
		n = n << 8
		n |= int(x)
	}
	return n
}
func readInt(r io.Reader, n uint) (int, error) {
	b, err := readBytes(r, n)
	if err != nil {
		return 0, err
	}
	return getInt(b), nil
}

func readUint(r io.Reader, n uint) (uint, error) {
	x, err := readInt(r, n)
	if err != nil {
		return 0, err
	}
	return uint(x), nil
}

// readBytesMaxUpfront is the max up-front allocation allowed
const readBytesMaxUpfront = 10 << 20 // 10MB

func readBytes(r io.Reader, n uint) ([]byte, error) {
	if n > readBytesMaxUpfront {
		b := &bytes.Buffer{}
		if _, err := io.CopyN(b, r, int64(n)); err != nil {
			return nil, err
		}
		return b.Bytes(), nil
	}

	b := make([]byte, n)
	_, err := io.ReadFull(r, b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func readString(r io.Reader, n uint) (string, error) {
	b, err := readBytes(r, n)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
func readUint32LittleEndian(r io.Reader) (uint32, error) {
	b, err := readBytes(r, 4)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint32(b), nil
}
func encodeUint32(n uint32) []byte {
	buf := bytes.NewBuffer([]byte{})
	if err := binary.Write(buf, binary.BigEndian, n); err != nil {
		panic(err)
	}
	return buf.Bytes()
}

func fileTypesContains(v FileType, a []FileType) bool {
	for _, i := range a {
		if i == v {
			return true
		}
	}
	return false
}
