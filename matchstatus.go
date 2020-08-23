package main

type MatchStatus struct {
	All       bool
	Failed    bool
	Args      []bool
	ActiveArg int
}

func NewMatchStatus(nArgs int) *MatchStatus {
	return &MatchStatus{
		Args: make([]bool,nArgs),
	}
}

func (s *MatchStatus) Reset() {
	s.All = false
	s.Failed = false
	for i := range s.Args {
		s.Args[i] = false
	}
	s.ActiveArg = 0
}

func (s *MatchStatus) Process() {
	switch {
	case s.Failed:
		s.All = false
	case s.All:
		// do nothing, s.All is already set	
	case s.ActiveArg != len(s.Args):
		// this means the amount of args the pattern was expecting
		// was not the amout of args the pattern got,
		// so the match fails
		s.All = false
	default:
		s.All = true
		for i:=0;i<s.ActiveArg;i++ {
			if s.Args[i] == false {
				s.All = false
			}
		}
	}
}

func (s *MatchStatus) MatchArg() {
	s.Args[s.ActiveArg] = true
}
