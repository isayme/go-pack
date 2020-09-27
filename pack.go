package pack

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Unpack read and unpack from io reader
func Unpack(r io.Reader) ([]byte, error) {
	var size uint32
	err := binary.Read(r, binary.BigEndian, &size)
	if err != nil {
		return nil, fmt.Errorf("unpack fail, read size: %w", err)
	}

	buf := make([]byte, size)
	_, err = io.ReadFull(r, buf)
	if err != nil {
		return nil, fmt.Errorf("unpack fail, read data: %w", err)
	}

	return buf, nil
}

// Pack pack and write to io writer
func Pack(w io.Writer, data []byte) error {
	size := uint32(len(data))
	err := binary.Write(w, binary.BigEndian, size)
	if err != nil {
		return fmt.Errorf("pack fail, write size: %w", err)
	}

	_, err = w.Write(data)
	if err != nil {
		return fmt.Errorf("pack fail, write data: %w", err)
	}

	return nil
}
