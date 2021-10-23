package save

import (
	"fmt"
	"io"
)

// ParseHeader will only parse the header of the save file and return it.
func ParseHeader(r io.Reader) (*Header, error) {
	p, err := newParser(r)
	if err != nil {
		return nil, err
	}

	return p.parseHeader()
}

type Header struct {
	HeaderVersion       int32  `json:"headerVersion"`
	SaveVersion         int32  `json:"saveVersion"`
	BuildVersion        int32  `json:"buildVersion"`
	MapName             string `json:"mapName"`
	MapOptions          string `json:"mapOptions"`
	SessionName         string `json:"sessionName"`
	PlayTime            int32  `json:"playTime"`
	SaveDate            int64  `json:"saveDate"`
	SessionVisibility   byte   `json:"sessionVisibility"`
	EditorObjectVersion int32  `json:"editorObjectVersion"`
	ModMetadata         string `json:"modMetadata"`
	ModFlags            int32  `json:"modFlags"`
}

func (p *parser) parseHeader() (*Header, error) {
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

	// TODO: Add Get method to return map[string]string
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

	// TODO: Constants for different visibility states
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

func (p *parser) serializeHeader(h *Header) error {
	err := p.writeInt32(h.HeaderVersion)
	if err != nil {
		return err
	}

	err = p.writeInt32(h.SaveVersion)
	if err != nil {
		return err
	}

	err = p.writeInt32(h.BuildVersion)
	if err != nil {
		return err
	}

	err = p.writeString(h.MapName)
	if err != nil {
		return err
	}

	err = p.writeString(h.MapOptions)
	if err != nil {
		return err
	}

	err = p.writeString(h.SessionName)
	if err != nil {
		return err
	}

	err = p.writeInt32(h.PlayTime)
	if err != nil {
		return err
	}

	err = p.writeInt64(h.SaveDate)
	if err != nil {
		return err
	}

	err = p.writeByte(h.SessionVisibility)
	if err != nil {
		return err
	}

	err = p.writeInt32(h.EditorObjectVersion)
	if err != nil {
		return err
	}

	err = p.writeString(h.ModMetadata)
	if err != nil {
		return err
	}

	err = p.writeInt32(h.ModFlags)
	if err != nil {
		return err
	}

	return nil
}
