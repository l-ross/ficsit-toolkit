package save

import (
	"encoding/binary"
	"fmt"
	"io"
)

//
// Read
//

func (p *Parser) readInt8() (int8, error) {
	var v int8
	err := binary.Read(p.body, binary.LittleEndian, &v)
	return v, err
}

func (p *Parser) readInt32() (int32, error) {
	var v int32
	err := binary.Read(p.body, binary.LittleEndian, &v)
	return v, err
}

func (p *Parser) readInt32Array(l int) ([]int32, error) {
	v := make([]int32, l)
	err := binary.Read(p.body, binary.LittleEndian, &v)
	return v, err
}

func (p *Parser) readInt64() (int64, error) {
	var v int64
	err := binary.Read(p.body, binary.LittleEndian, &v)
	return v, err
}

func (p *Parser) readFloat32() (float32, error) {
	var v float32
	err := binary.Read(p.body, binary.LittleEndian, &v)
	return v, err
}

func (p *Parser) readFloat64() (float64, error) {
	var v float64
	err := binary.Read(p.body, binary.LittleEndian, &v)
	return v, err
}

func (p *Parser) readFloat32Array(l int) ([]float32, error) {
	v := make([]float32, l)
	err := binary.Read(p.body, binary.LittleEndian, &v)
	return v, err
}

func (p *Parser) readByte() (byte, error) {
	v, err := p.readBytes(1)
	if err != nil {
		return 0, err
	}
	return v[0], nil
}

func (p *Parser) readBytes(l int32) ([]byte, error) {
	v := make([]byte, l)
	err := binary.Read(p.body, binary.LittleEndian, &v)
	return v, err
}

func (p *Parser) readString() (string, error) {
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

func (p *Parser) readBool() (bool, error) {
	b, err := p.readByte()
	if err != nil {
		return false, nil
	}

	if b == 0 {
		return false, nil
	}
	return true, nil
}

func (p *Parser) nextByteIsNull() error {
	b, err := p.readByte()
	if err != nil {
		return err
	}
	if b != 0 {
		return fmt.Errorf("expected byte to be null")
	}

	return nil
}

func (p *Parser) skipBytes(l int64) error {
	_, err := p.body.Seek(l, io.SeekCurrent)
	return err
}

// Reads a data chunk and returns the bytes with the length of the
// chunk prepended.
// Resets p.body back to where the function started reading from.
// Useful when debugging a data chunk.
func (p *Parser) debugReadDataChunk(l int32) ([]byte, error) {
	pos := p.body.Index

	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, uint32(l))

	b2, err := p.readBytes(l)
	if err != nil {
		return nil, err
	}

	b = append(b, b2...)

	_, err = p.body.Seek(pos, io.SeekStart)
	if err != nil {
		return nil, err
	}

	return b, nil

}

// measure returns a function that will return the difference between the index when measure was
// called and when the returned function was called.
func (p *Parser) measure() func() int32 {
	startPos := p.body.Index
	return func() int32 {
		return int32(p.body.Index - startPos)
	}
}

//
// Write
//

func (s *Serializer) writeByte(b byte) error {
	_, err := s.body.Write([]byte{b})
	return err
}

func (s *Serializer) writeBytes(b []byte) error {
	_, err := s.body.Write(b)
	return err
}

func (s *Serializer) writeInt8(i int8) error {
	return binary.Write(s.body, binary.LittleEndian, i)
}

func (s *Serializer) writeInt32(i int32) error {
	return binary.Write(s.body, binary.LittleEndian, i)
}

func (s *Serializer) writeInt32Array(i []int32) error {
	return binary.Write(s.body, binary.LittleEndian, i)
}

func (s *Serializer) writeInt64(i int64) error {
	return binary.Write(s.body, binary.LittleEndian, i)
}

func (s *Serializer) writeFloat32(f float32) error {
	return binary.Write(s.body, binary.LittleEndian, f)
}

func (s *Serializer) writeFloat32Array(f []float32) error {
	return binary.Write(s.body, binary.LittleEndian, f)
}

func (s *Serializer) writeFloat64(f float64) error {
	return binary.Write(s.body, binary.LittleEndian, f)
}

func (s *Serializer) writeBool(b bool) error {
	if b {
		return binary.Write(s.body, binary.LittleEndian, byte(0x01))
	}
	return binary.Write(s.body, binary.LittleEndian, byte(0x00))
}

func (s *Serializer) writeNull() error {
	return binary.Write(s.body, binary.LittleEndian, byte(0x00))
}

func (s *Serializer) writeNulls(len int32) error {
	for i := int32(0); i < len; i++ {
		err := s.writeNull()
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Serializer) writeString(str string) error {
	if len(str) == 0 {
		return s.writeInt32(0)
	}

	// Add null termination.
	str += "\x00"

	err := s.writeInt32(int32(len(str)))
	if err != nil {
		return err
	}

	return binary.Write(s.body, binary.LittleEndian, []byte(str))
}

func (s *Serializer) writeNoneProp() error {
	return s.writeString("None")
}

// writeLen writes the int32 at the specified index.
// Resets its position to its original location after writing.
func (s *Serializer) writeLen(i int32, idx int64) error {
	currentPos := s.body.Index

	_, err := s.body.Seek(idx, io.SeekStart)
	if err != nil {
		return err
	}

	err = s.writeInt32(i)
	if err != nil {
		return err
	}

	_, err = s.body.Seek(currentPos, io.SeekStart)
	if err != nil {
		return err
	}

	return nil
}

// measure returns a function that will return the difference between the index when measure was
// called and when the returned function was called.
func (s *Serializer) measure() func() int32 {
	startPos := s.body.Index
	return func() int32 {
		return int32(s.body.Index - startPos)
	}
}
