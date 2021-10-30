package save

type serializerOptions struct {
}

type SerializerOption func(*serializerOptions) error

func defaultSerializerOptions() serializerOptions {
	return serializerOptions{}
}
