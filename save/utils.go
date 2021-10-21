package save

import (
	"encoding/binary"
	"fmt"
	"io"
)

//
// Read
//

func (p *parser) readInt8() (int8, error) {
	var v int8
	err := binary.Read(p.body, binary.LittleEndian, &v)
	return v, err
}

func (p *parser) readInt32() (int32, error) {
	var v int32
	err := binary.Read(p.body, binary.LittleEndian, &v)
	return v, err
}

func (p *parser) readInt32Array(l int) ([]int32, error) {
	v := make([]int32, l)
	err := binary.Read(p.body, binary.LittleEndian, &v)
	return v, err
}

func (p *parser) readInt64() (int64, error) {
	var v int64
	err := binary.Read(p.body, binary.LittleEndian, &v)
	return v, err
}

func (p *parser) readFloat32() (float32, error) {
	var v float32
	err := binary.Read(p.body, binary.LittleEndian, &v)
	return v, err
}

func (p *parser) readFloat64() (float64, error) {
	var v float64
	err := binary.Read(p.body, binary.LittleEndian, &v)
	return v, err
}

func (p *parser) readFloat32Array(l int) ([]float32, error) {
	v := make([]float32, l)
	err := binary.Read(p.body, binary.LittleEndian, &v)
	return v, err
}

func (p *parser) readByte() (byte, error) {
	v, err := p.readBytes(1)
	if err != nil {
		return 0, err
	}
	return v[0], nil
}

func (p *parser) readBytes(l int32) ([]byte, error) {
	v := make([]byte, l)
	err := binary.Read(p.body, binary.LittleEndian, &v)
	return v, err
}

func (p *parser) readString() (string, error) {
	l, err := p.readInt32()
	if err != nil {
		return "", err
	}

	if l == 0 {
		return "", nil
	}

	v := make([]byte, l)
	read, err := p.body.Read(v)
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

func (p *parser) readBool() (bool, error) {
	b, err := p.readByte()
	if err != nil {
		return false, nil
	}

	if b == 0 {
		return false, nil
	}
	return true, nil
}

func (p *parser) nextByteIsNull() error {
	b, err := p.readByte()
	if err != nil {
		return err
	}
	if b != 0 {
		return fmt.Errorf("expected byte to be null")
	}

	return nil
}

func (p *parser) skipBytes(l int64) error {
	_, err := p.body.Seek(l, io.SeekCurrent)
	return err
}

//
// Write
//

func (p *parser) writeByte(b byte) error {
	_, err := p.body.Write([]byte{b})
	return err
}

func (p *parser) writeBytes(b []byte) error {
	_, err := p.body.Write(b)
	return err
}

func (p *parser) writeInt8(i int8) error {
	return binary.Write(p.body, binary.LittleEndian, i)
}

func (p *parser) writeInt32(i int32) error {
	return binary.Write(p.body, binary.LittleEndian, i)
}

func (p *parser) writeInt32Array(i []int32) error {
	return binary.Write(p.body, binary.LittleEndian, i)
}

func (p *parser) writeInt64(i int64) error {
	return binary.Write(p.body, binary.LittleEndian, i)
}

func (p *parser) writeFloat32(f float32) error {
	return binary.Write(p.body, binary.LittleEndian, f)
}

func (p *parser) writeFloat32Array(f []float32) error {
	return binary.Write(p.body, binary.LittleEndian, f)
}

func (p *parser) writeFloat64(f float64) error {
	return binary.Write(p.body, binary.LittleEndian, f)
}

func (p *parser) writeBool(b bool) error {
	if b {
		return binary.Write(p.body, binary.LittleEndian, byte(0x01))
	}
	return binary.Write(p.body, binary.LittleEndian, byte(0x00))
}

func (p *parser) writeNull() error {
	return binary.Write(p.body, binary.LittleEndian, byte(0x00))
}

func (p *parser) writeNulls(len int32) error {
	for i := int32(0); i < len; i++ {
		err := p.writeNull()
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *parser) writeString(str string) error {
	if len(str) == 0 {
		return p.writeInt32(0)
	}

	// Add null termination.
	str += "\x00"

	err := p.writeInt32(int32(len(str)))
	if err != nil {
		return err
	}

	return binary.Write(p.body, binary.LittleEndian, []byte(str))
}

func (p *parser) writeNoneProp() error {
	return p.writeString("None")
}

// writeLen writes the int32 at the specified index.
// Resets its position to its original location after writing.
func (p *parser) writeLen(i int32, idx int64) error {
	currentPos := p.body.Index

	_, err := p.body.Seek(idx, io.SeekStart)
	if err != nil {
		return err
	}

	err = p.writeInt32(i)
	if err != nil {
		return err
	}

	_, err = p.body.Seek(currentPos, io.SeekStart)
	if err != nil {
		return err
	}

	return nil
}

// measure returns a function that will return the difference between the index when measure was
// called and when the returned function was called.
func (p *parser) measure() func() int32 {
	startPos := p.body.Index
	return func() int32 {
		return int32(p.body.Index - startPos)
	}
}
