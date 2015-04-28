package main

// START OMIT
type Server struct {
	Do       chan Request
	internal map[string]string
}

func (s *Server) do(r Request) error {
	secret := s.internal[r.Target]
	// do something, and maybe update the internal map
	s.internal[r.Target] = "better " + secret
	return nil
}

func (s *Server) Start() {
	for r := range s.Do {
		// still not really concurrent
		r.Err <- s.do(r)
	}
}

type Request struct {
	Target string
	Action string
	Err    chan error
}

// END OMIT

func main() {
}
