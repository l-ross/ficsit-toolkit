package data

import (
	"encoding/binary"
	"io"
	"strings"
	"unicode"
)

func (d *Data) Write(p []byte) (int, error) {
	return d.sws.Write(p)
}

func (d *Data) WriteByte(b byte) error {
	_, err := d.sws.Write([]byte{b})
	return err
}

func (d *Data) WriteBytes(b []byte) error {
	_, err := d.sws.Write(b)
	return err
}

func (d *Data) WriteInt8(i int8) error {
	return binary.Write(d.sws, binary.LittleEndian, i)
}

func (d *Data) WriteInt32(i int32) error {
	return binary.Write(d.sws, binary.LittleEndian, i)
}

func (d *Data) WriteInt32Array(i []int32) error {
	return binary.Write(d.sws, binary.LittleEndian, i)
}

func (d *Data) WriteInt64(i int64) error {
	return binary.Write(d.sws, binary.LittleEndian, i)
}

func (d *Data) WriteFloat32(f float32) error {
	return binary.Write(d.sws, binary.LittleEndian, f)
}

func (d *Data) WriteFloat32Array(f []float32) error {
	return binary.Write(d.sws, binary.LittleEndian, f)
}

func (d *Data) WriteFloat64(f float64) error {
	return binary.Write(d.sws, binary.LittleEndian, f)
}

func (d *Data) WriteBool(b bool) error {
	if b {
		return binary.Write(d.sws, binary.LittleEndian, byte(0x01))
	}
	return binary.Write(d.sws, binary.LittleEndian, byte(0x00))
}

func (d *Data) WriteNull() error {
	return binary.Write(d.sws, binary.LittleEndian, byte(0x00))
}

func (d *Data) WriteNulls(len int32) error {
	for i := int32(0); i < len; i++ {
		err := d.WriteNull()
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *Data) WriteString(str string) error {
	if len(str) == 0 {
		return d.WriteInt32(0)
	}

	// Add null termination.
	str += "\x00"

	strLen := int32(strings.Count(str, "")) - 1

	// If the string isn't just ASCII characters then adjust the length and encode
	// the string.
	if !isASCII(str) {
		strLen = strLen * -1

		var err error
		str, err = utf16Encoder.String(str)
		if err != nil {
			return err
		}
	}

	err := d.WriteInt32(strLen)
	if err != nil {
		return err
	}

	return binary.Write(d.sws, binary.LittleEndian, []byte(str))
}

func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func (d *Data) WriteNoneProp() error {
	return d.WriteString("None")
}

// WriteLen writes the int32 at the specified index and then resets the Index
// of Data back to its original position.
//
// Useful in situations where we don't know what the length of something will be until
// we have written it. We can write a placeholder length, write the data and then update the length.
func (d *Data) WriteLen(i int32, idx int64) error {
	currentPos := d.sws.Index

	_, err := d.sws.Seek(idx, io.SeekStart)
	if err != nil {
		return err
	}

	err = d.WriteInt32(i)
	if err != nil {
		return err
	}

	_, err = d.sws.Seek(currentPos, io.SeekStart)
	if err != nil {
		return err
	}

	return nil
}
