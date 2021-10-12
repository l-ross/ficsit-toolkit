package satisfactorysave

import (
	"encoding/binary"
	"fmt"
	"io"
)

func readInt8(r io.Reader) (int8, error) {
	var v int8
	err := binary.Read(r, binary.LittleEndian, &v)
	return v, err
}

func readInt32(r io.Reader) (int32, error) {
	var v int32
	err := binary.Read(r, binary.LittleEndian, &v)
	return v, err
}

func readInt64(r io.Reader) (int64, error) {
	var v int64
	err := binary.Read(r, binary.LittleEndian, &v)
	return v, err
}
func readFloat32(r io.Reader) (float32, error) {
	var v float32
	err := binary.Read(r, binary.LittleEndian, &v)
	return v, err
}

func readFloat64(r io.Reader) (float64, error) {
	var v float64
	err := binary.Read(r, binary.LittleEndian, &v)
	return v, err
}

func readFloat32Array(r io.Reader, l int) ([]float32, error) {
	v := make([]float32, l)

	for i := 0; i < l; i++ {
		f, err := readFloat32(r)
		if err != nil {
			return nil, err
		}

		v[i] = f
	}

	return v, nil
}

func readByte(r io.Reader) (byte, error) {
	v, err := readBytes(r, 1)
	if err != nil {
		return 0, err
	}
	return v[0], nil
}

func readBytes(r io.Reader, l int32) ([]byte, error) {
	v := make([]byte, l)
	read, err := r.Read(v)
	if err != nil {
		return nil, err
	}
	if read != int(l) {
		return nil, fmt.Errorf("expected to read %d but read %d", l, read)
	}
	return v, nil
}

func readString(r io.Reader) (string, error) {
	l, err := readInt32(r)
	if err != nil {
		return "", err
	}

	if l == 0 {
		return "", nil
	}

	v := make([]byte, l)
	read, err := r.Read(v)
	if err != nil {
		return "", err
	}
	if read != int(l) {
		return "", fmt.Errorf("expected to read %d but only read %d", l, read)
	}

	// Drop the null terminator
	v = v[:l-1]
	return string(v), nil
}

func readBool(r io.Reader) (bool, error) {
	b, err := readByte(r)
	if err != nil {
		return false, nil
	}

	err = nextByteIsNull(r)
	if err != nil {
		return false, err
	}

	if b == 0 {
		return false, nil
	}
	return true, nil
}

func nextByteIsNull(r io.Reader) error {
	b, err := readByte(r)
	if err != nil {
		return err
	}
	if b != 0 {
		return fmt.Errorf("expected byte to be null")
	}
	return nil
}
