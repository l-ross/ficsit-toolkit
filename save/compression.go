package save

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/l-ross/ficsit-toolkit/save/data"
)

const (
	maximumChunkSize = int64(131072)
	packageFileTag   = int64(2653586369)
)

func (p *parser) decompressBody() (*data.Data, error) {
	chunks := make([]byte, 0)

	for {
		ch, err := p.readChunkHeader()
		if err != nil {
			return nil, err
		}
		if ch == nil {
			// No more chunks.
			break
		}

		chunk, err := p.readChunk(ch)
		if err != nil {
			return nil, err
		}

		chunks = append(chunks, chunk...)
	}

	return data.NewFromBytes(chunks), nil
}

type chunkHeader struct {
	packageFileTag     int64
	maximumChunkSize   int64
	compressedLength   int64
	uncompressedLength int64
}

func (p *parser) readChunkHeader() (*chunkHeader, error) {
	ch := &chunkHeader{}

	var err error

	ch.packageFileTag, err = p.ReadInt64()
	if err != nil {
		if err == io.EOF {
			return nil, nil
		}

		return nil, err
	}

	ch.maximumChunkSize, err = p.ReadInt64()
	if err != nil {
		return nil, err
	}

	ch.compressedLength, err = p.ReadInt64()
	if err != nil {
		return nil, err
	}

	ch.uncompressedLength, err = p.ReadInt64()
	if err != nil {
		return nil, err
	}

	// Skip the next 16 bytes as they are seemingly just a repeat the compressed
	// and uncompressed lengths
	err = p.SkipBytes(16)
	if err != nil {
		return nil, err
	}

	return ch, nil
}

// readChunk will read the next chunk and return the decompressed data.
func (p *parser) readChunk(ch *chunkHeader) ([]byte, error) {
	compressed := make([]byte, ch.compressedLength)

	read, err := p.Read(compressed)
	if err != nil {
		return nil, err
	}
	if int64(read) != ch.compressedLength {
		return nil, fmt.Errorf("expected to read %d but got %d", ch.compressedLength, read)
	}

	zr, err := zlib.NewReader(bytes.NewBuffer(compressed))
	if err != nil {
		return nil, err
	}
	defer zr.Close()

	uncompressed, err := ioutil.ReadAll(zr)
	if err != nil {
		return nil, err
	}
	if int64(len(uncompressed)) != ch.uncompressedLength {
		return nil, fmt.Errorf("expected uncompressed lenght to be %d but was %d", ch.uncompressedLength, len(uncompressed))
	}

	return uncompressed, nil
}

func (s *serializer) compressBody(b []byte) error {
	done := false

	for !done {
		var chunk []byte
		if int64(len(b)) < maximumChunkSize {
			// Last chunk
			chunk = b
			done = true
		} else {
			chunk, b = b[:maximumChunkSize], b[maximumChunkSize:]
		}

		var cb bytes.Buffer
		zw := zlib.NewWriter(&cb)

		_, err := zw.Write(chunk)
		if err != nil {
			return err
		}

		err = zw.Close()
		if err != nil {
			return err
		}

		ch := &chunkHeader{
			packageFileTag:     packageFileTag,
			maximumChunkSize:   maximumChunkSize,
			compressedLength:   int64(len(cb.Bytes())),
			uncompressedLength: int64(len(chunk)),
		}

		err = s.writeChunkHeader(ch)
		if err != nil {
			return err
		}

		err = s.WriteBytes(cb.Bytes())
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *serializer) writeChunkHeader(ch *chunkHeader) error {
	err := s.WriteInt64(ch.packageFileTag)
	if err != nil {
		return err
	}

	err = s.WriteInt64(ch.maximumChunkSize)
	if err != nil {
		return err
	}

	err = s.WriteInt64(ch.compressedLength)
	if err != nil {
		return err
	}

	err = s.WriteInt64(ch.uncompressedLength)
	if err != nil {
		return err
	}

	// Duplicate the compressed and uncompressed lengths.

	err = s.WriteInt64(ch.compressedLength)
	if err != nil {
		return err
	}

	err = s.WriteInt64(ch.uncompressedLength)
	if err != nil {
		return err
	}

	return nil
}
