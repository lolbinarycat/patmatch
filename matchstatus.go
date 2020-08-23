package main

type MatchStatus struct {
	All  bool
	Args []bool
}

func NewMatchStatus(nArgs int) *MatchStatus {
	return &MatchStatus{
		Args: make([]bool,nArgs),
	}
}

func (s *MatchStatus) Falsify() {
	s.All = false
	for i := range s.Args {
		s.Args[i] = false
	}
}
