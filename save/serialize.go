package save

import "io"

func (s *Save) Serialize(w io.Writer) error {
	s.w = w

	return nil
}
