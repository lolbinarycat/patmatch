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
	logln("MatchStatus reset")
}

func (s *MatchStatus) Process() {
	logln("processing mstat")
	switch {
	case s.Failed:
		s.All = false
	case s.All:
		// do nothing, s.All is already set	
	case s.ActiveArg != len(s.Args):
		// this means the amount of args the pattern was expecting
		// was not the amout of args the pattern got,
		// so the match fails
		logln("wrong number of args, needed",len(s.Args),"got",s.ActiveArg)
		s.All = false
	default:
		s.All = true
		for i:=0;i<s.ActiveArg;i++ {
			if s.Args[i] == false {
				logln("arg",i,"did not match")
				s.All = false
			}
		}
	}
	logln("processed mstat")
}

func (s *MatchStatus) MatchArg() {
	s.Args[s.ActiveArg] = true
}
