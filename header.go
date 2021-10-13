package satisfactorysave

import (
	"fmt"
)

// ParseHeader will only parse the header of the save file and return it.
func (p *Parser) ParseHeader() (*Header, error) {
	h := &Header{}

	var err error

	h.HeaderVersion, err = p.readInt32()
	if err != nil {
		return nil, err
	}

	// TODO: Better version handling.
	if h.HeaderVersion != 8 {
		return nil, fmt.Errorf("only support save header version 8, got %d", h.HeaderVersion)
	}

	h.SaveVersion, err = p.readInt32()
	if err != nil {
		return nil, err
	}

	h.BuildVersion, err = p.readInt32()
	if err != nil {
		return nil, err
	}

	h.MapName, err = p.readString()
	if err != nil {
		return nil, err
	}

	// TODO: Convert to a map?
	h.MapOptions, err = p.readString()
	if err != nil {
		return nil, err
	}

	h.SessionName, err = p.readString()
	if err != nil {
		return nil, err
	}

	h.PlayTime, err = p.readInt32()
	if err != nil {
		return nil, err
	}

	h.SaveDate, err = p.readInt64()
	if err != nil {
		return nil, err
	}

	h.SessionVisibility, err = p.readByte()
	if err != nil {
		return nil, err
	}

	h.EditorObjectVersion, err = p.readInt32()
	if err != nil {
		return nil, err
	}

	h.ModMetadata, err = p.readString()
	if err != nil {
		return nil, err
	}

	h.ModFlags, err = p.readInt32()
	if err != nil {
		return nil, err
	}

	return h, nil
}
