package satisfactorysave

import "fmt"

type StructType string

const ()

//
// ArbitraryStruct
//

type ArbitraryStruct struct {
	Properties []*Property
}

func (s *StructPropertyValue) GetArbitraryStruct() (*ArbitraryStruct, error) {
	if v, ok := s.Value.(*ArbitraryStruct); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong type %s", s.Type)
}

func (s *ArbitraryStruct) parse(p *Parser) error {
	for {
		prop, err := p.parseProperty()
		if err != nil {
			return err
		}
		if prop == nil {
			break
		}

		s.Properties = append(s.Properties, prop)
	}

	return nil
}
