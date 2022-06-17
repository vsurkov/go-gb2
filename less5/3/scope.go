package main

type Scope struct {
	writePercent float32
	scope        float32
}

func NewScope(writePercent float32, scope float32) *Scope {
	return &Scope{
		writePercent: writePercent,
		scope:        scope,
	}
}

func (s *Scope) GetWriteScope() int {
	return int(s.scope * (s.writePercent / 100))
}
func (s *Scope) GetReadScope() int {
	return int(s.scope * (1 - (s.writePercent / 100)))
}
