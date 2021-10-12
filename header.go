package satisfactorysave

import (
	"fmt"
)

// ParseHeader will only parse the header of the save file and return it.
func (p *Parser) ParseHeader() (*Header, error) {
	h := &Header{}

	var err error

	h.HeaderVersion, err = readInt32(p.r)
	if err != nil {
		return nil, err
	}

	// TODO: Better version handling.
	if h.HeaderVersion != 8 {
		return nil, fmt.Errorf("only support save header version 8, got %d", h.HeaderVersion)
	}

	h.SaveVersion, err = readInt32(p.r)
	if err != nil {
		return nil, err
	}

	h.BuildVersion, err = readInt32(p.r)
	if err != nil {
		return nil, err
	}

	h.MapName, err = readString(p.r)
	if err != nil {
		return nil, err
	}

	// TODO: Convert to a map?
	h.MapOptions, err = readString(p.r)
	if err != nil {
		return nil, err
	}

	h.SessionName, err = readString(p.r)
	if err != nil {
		return nil, err
	}

	h.PlayTime, err = readInt32(p.r)
	if err != nil {
		return nil, err
	}

	h.SaveDate, err = readInt64(p.r)
	if err != nil {
		return nil, err
	}

	h.SessionVisibility, err = readByte(p.r)
	if err != nil {
		return nil, err
	}

	h.EditorObjectVersion, err = readInt32(p.r)
	if err != nil {
		return nil, err
	}

	h.ModMetadata, err = readString(p.r)
	if err != nil {
		return nil, err
	}

	h.ModFlags, err = readInt32(p.r)
	if err != nil {
		return nil, err
	}

	return h, nil
}
