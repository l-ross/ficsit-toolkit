package save

type parserOptions struct {
}

type ParserOption func(*parserOptions) error

func defaultParserOptions() parserOptions {
	return parserOptions{}
}
