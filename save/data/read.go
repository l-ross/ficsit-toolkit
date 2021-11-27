package data

import (
	"encoding/binary"
	"fmt"
	"io"
)

func (d *Data) Read(p []byte) (n int, err error) {
	return d.sws.Read(p)
}

func (d *Data) ReadInt8() (int8, error) {
	var v int8
	err := binary.Read(d.sws, binary.LittleEndian, &v)
	return v, err
}

func (d *Data) ReadInt32() (int32, error) {
	var v int32
	err := binary.Read(d.sws, binary.LittleEndian, &v)
	return v, err
}

func (d *Data) ReadInt32Array(l int) ([]int32, error) {
	v := make([]int32, l)
	err := binary.Read(d.sws, binary.LittleEndian, &v)
	return v, err
}

func (d *Data) ReadInt64() (int64, error) {
	var v int64
	err := binary.Read(d.sws, binary.LittleEndian, &v)
	return v, err
}

func (d *Data) ReadFloat32() (float32, error) {
	var v float32
	err := binary.Read(d.sws, binary.LittleEndian, &v)
	return v, err
}

func (d *Data) ReadFloat64() (float64, error) {
	var v float64
	err := binary.Read(d.sws, binary.LittleEndian, &v)
	return v, err
}

func (d *Data) ReadFloat32Array(l int) ([]float32, error) {
	v := make([]float32, l)
	err := binary.Read(d.sws, binary.LittleEndian, &v)
	return v, err
}

func (d *Data) ReadByte() (byte, error) {
	v, err := d.ReadBytes(1)
	if err != nil {
		return 0, err
	}
	return v[0], nil
}

func (d *Data) ReadBytes(l int32) ([]byte, error) {
	v := make([]byte, l)
	err := binary.Read(d.sws, binary.LittleEndian, &v)
	return v, err
}

func (d *Data) ReadString() (string, error) {
	l, err := d.ReadInt32()
	if err != nil {
		return "", err
	}

	if l == 0 {
		return "", nil
	}

	// If the length is negative then it's UTF16 encoded.
	isUTF16 := false
	if l < 0 {
		isUTF16 = true
		l = l * -2
	}

	// Read all the bytes that make up the string.
	b := make([]byte, l)
	read, err := d.sws.Read(b)
	if err != nil {
		return "", err
	}
	if read != int(l) {
		return "", fmt.Errorf("expected to read %d but only read %d", l, read)
	}

	// Convert bytes to string
	v := ""

	if isUTF16 {
		b, err := utf16Decoder.Bytes(b)
		if err != nil {
			return "", err
		}
		v = string(b)
	} else {
		v = string(b)
	}

	// Remove null terminator
	v = v[:len(v)-1]

	return v, nil
}

func (d *Data) ReadBool() (bool, error) {
	b, err := d.ReadByte()
	if err != nil {
		return false, nil
	}

	if b == 0 {
		return false, nil
	}
	return true, nil
}

// NextByteIsNull will return an error if the next byte is not a null.
func (d *Data) NextByteIsNull() error {
	b, err := d.ReadByte()
	if err != nil {
		return err
	}
	if b != 0 {
		return fmt.Errorf("expected byte to be null")
	}

	return nil
}

func (d *Data) SkipBytes(l int64) error {
	_, err := d.sws.Seek(l, io.SeekCurrent)
	return err
}

// DebugReadDataChunk reads a data chunk and returns the bytes with the length of the
// chunk prepended.
//
// Once reading has completed the index is wound back to the starting position before
// the read. Useful in some debugging scenarios.
func (d *Data) DebugReadDataChunk(l int32) ([]byte, error) {
	pos := d.sws.Index

	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, uint32(l))

	b2, err := d.ReadBytes(l)
	if err != nil {
		return nil, err
	}

	b = append(b, b2...)

	_, err = d.sws.Seek(pos, io.SeekStart)
	if err != nil {
		return nil, err
	}

	return b, nil

}
