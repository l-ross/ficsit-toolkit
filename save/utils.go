package save

import (
	"encoding/binary"
	"fmt"
)

func (p *Parser) resetCounter() {
	p.counter = 0
}

func (p *Parser) offset() int64 {
	return p.buf.Size() - int64(p.buf.Len())
}

func (p *Parser) readInt8() (int8, error) {
	var v int8
	err := binary.Read(p.buf, binary.LittleEndian, &v)
	p.counter += 1
	return v, err
}

func (p *Parser) readInt32() (int32, error) {
	var v int32
	err := binary.Read(p.buf, binary.LittleEndian, &v)
	p.counter += 4
	return v, err
}

func (p *Parser) readInt32Array(l int) ([]int32, error) {
	v := make([]int32, l)

	for i := 0; i < l; i++ {
		vv, err := p.readInt32()
		if err != nil {
			return nil, err
		}

		v[i] = vv
	}

	return v, nil
}

func (p *Parser) readInt64() (int64, error) {
	var v int64
	err := binary.Read(p.buf, binary.LittleEndian, &v)
	p.counter += 8
	return v, err
}

func (p *Parser) readFloat32() (float32, error) {
	var v float32
	err := binary.Read(p.buf, binary.LittleEndian, &v)
	p.counter += 4
	return v, err
}

func (p *Parser) readFloat64() (float64, error) {
	var v float64
	err := binary.Read(p.buf, binary.LittleEndian, &v)
	p.counter += 8
	return v, err
}

func (p *Parser) readFloat32Array(l int) ([]float32, error) {
	v := make([]float32, l)

	for i := 0; i < l; i++ {
		f, err := p.readFloat32()
		if err != nil {
			return nil, err
		}

		v[i] = f
	}

	p.counter += int32(l)

	return v, nil
}

func (p *Parser) readByte() (byte, error) {
	v, err := p.readBytes(1)
	if err != nil {
		return 0, err
	}
	// Don't increment counter.
	return v[0], nil
}

func (p *Parser) readBytes(l int32) ([]byte, error) {
	v := make([]byte, l)

	totalRead := int32(0)

	for totalRead < l {
		read, err := p.buf.Read(v)
		if err != nil {
			return nil, err
		}

		totalRead += int32(read)
	}

	p.counter += l
	return v, nil
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
	read, err := p.buf.Read(v)
	if err != nil {
		return "", err
	}
	if read != int(l) {
		return "", fmt.Errorf("expected to read %d but only read %d", l, read)
	}

	p.counter += l

	// Drop the null terminator
	v = v[:l-1]
	return string(v), nil
}

func (p *Parser) readBool() (bool, error) {
	b, err := p.readByte()
	if err != nil {
		return false, nil
	}

	// Don't increment counter.

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

	// Don't increment counter.

	return nil
}
