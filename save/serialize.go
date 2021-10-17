package save

import "io"

func (s *Save) Serialize(w io.Writer) error {
	s.w = w

	err := s.serializeHeader()
	if err != nil {
		return err
	}

	return nil
}
