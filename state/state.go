package state

type State struct {
	Data map[string]interface{}
}

func (s *State) Update(key string, value interface{}) {
	s.Data[key] = value
}

func (s *State) Get(key string) (interface{}, bool) {
	value, exists := s.Data[key]
	return value, exists
}

func NewState() *State {
	return &State{
		Data: make(map[string]interface{}),
	}
}
