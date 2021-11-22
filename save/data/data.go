package data

import (
	"io"
	"io/ioutil"

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

func (d *Data) Len() int64 {
	return d.sws.Len()
}

// Measure returns a function that will return the difference between the index when measure was
// called and when the returned function was called.
func (d *Data) Measure() func() int32 {
	startPos := d.sws.Index
	return func() int32 {
		return int32(d.sws.Index - startPos)
	}
}
