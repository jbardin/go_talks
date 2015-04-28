package main

import "sync"

// START OMIT
type Server struct {
	sync.Mutex
	// don't touch!
	internal map[string]string
}

func (s *Server) Do(r Request) error {
	s.Lock()
	defer s.Unlock()

	secret := s.internal[r.Target]
	// do something, and maybe update the internal map
	s.internal[r.Target] = "better " + secret
	return nil
}

type Request struct {
	Target string
	Action string
}

// END OMIT

func main() {
}
