package data

import (
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"unicode"

	"github.com/ViRb3/slicewriteseek"
	unicodex "golang.org/x/text/encoding/unicode"
)

var (
	utf16Decoder = unicodex.UTF16(unicodex.LittleEndian, unicodex.IgnoreBOM).NewDecoder()
	utf16Encoder = unicodex.UTF16(unicodex.LittleEndian, unicodex.IgnoreBOM).NewEncoder()
)

type Data struct {
	sws *slicewriteseek.SliceWriteSeeker
}

func New() *Data {
	return &Data{
		sws: slicewriteseek.New(),
	}
}

func NewFromReader(r io.Reader) (*Data, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	d := &Data{
		sws: &slicewriteseek.SliceWriteSeeker{
			Buffer: b,
		},
	}

	return d, nil
}

func NewFromBytes(b []byte) *Data {
	return &Data{
		sws: &slicewriteseek.SliceWriteSeeker{
			Buffer: b,
		},
	}
}

func (d *Data) Index() int64 {
	return d.sws.Index
}

func (d *Data) Bytes() []byte {
	return d.sws.Buffer
}

//func (d *Data) SetIndex(i int64) {
//	d.sws.Index = i
//}

func (d *Data) Len() int64 {
	return d.sws.Len()
}

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

// Reads a data chunk and returns the bytes with the length of the
// chunk prepended.
// Resets p.sws back to where the function started reading from.
// Useful when debugging a data chunk.
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

// Measure returns a function that will return the difference between the index when measure was
// called and when the returned function was called.
func (d *Data) Measure() func() int32 {
	startPos := d.sws.Index
	return func() int32 {
		return int32(d.sws.Index - startPos)
	}
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

// writeLen writes the int32 at the specified index.
// Resets its position to its original location after writing.
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
