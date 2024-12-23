package server

type Store interface {
	Read(string) string
	Write(string) string
}

type InMemoryStore struct {
	data string
}

func (s *InMemoryStore) Read() string {
	return DeSerialize(s.data)
}

func (s *InMemoryStore) Write(data string) {
	s.data = Serialize(data)
}

// TODO: Add Serialization logic
func Serialize(data string) string {
	return data
}

func DeSerialize(data string) string {
	return data
}
