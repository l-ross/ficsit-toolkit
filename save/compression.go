package save

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"io/ioutil"
)

func (p *Parser) decompressBody() ([]byte, error) {
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

	return chunks, nil
}

type chunkHeader struct {
	packageFileTag     int64
	maximumChunkSize   int64
	compressedLength   int64
	uncompressedLength int64
}

func (p *Parser) readChunkHeader() (*chunkHeader, error) {
	ch := &chunkHeader{}

	var err error

	ch.packageFileTag, err = p.readInt64()
	if err != nil {
		if err == io.EOF {
			return nil, nil
		}

		return nil, err
	}

	ch.maximumChunkSize, err = p.readInt64()
	if err != nil {
		return nil, err
	}

	ch.compressedLength, err = p.readInt64()
	if err != nil {
		return nil, err
	}

	ch.uncompressedLength, err = p.readInt64()
	if err != nil {
		return nil, err
	}

	// For some reason the compressed and uncompressed lengths are repeated
	// TODO: Make this less crap
	_, err = p.readInt64()
	if err != nil {
		return nil, err
	}
	_, err = p.readInt64()
	if err != nil {
		return nil, err
	}

	return ch, nil
}

// readChunk will read the next chunk and return the decompressed data.
func (p *Parser) readChunk(ch *chunkHeader) ([]byte, error) {
	compressed := make([]byte, ch.compressedLength)

	read, err := p.buf.Read(compressed)
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
