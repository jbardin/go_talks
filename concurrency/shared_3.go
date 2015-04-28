package main

// START OMIT
type Server struct {
	Do   chan Request
	Send chan Request
	Done chan struct{}
	// don't touch!
	internal map[string]string
}

func (s *Server) Start() {
	for {
		select {
		case r := <-s.Do:
			go s.do(r)
		case r := <-s.Send:
			go s.send(r)
		case <-s.Done:
			return
		}
	}
}

// END OMIT

func (s *Server) do(r Request) {
	secret := s.internal[r.Target]
	// do something, and maybe update the internal map
	s.internal[r.Target] = "better " + secret
	r.Err <- nil
}

func (s *Server) send(r Request) {
	secret := s.internal[r.Target]
	// do something, and maybe update the internal map
	s.internal[r.Target] = "better " + secret
	r.Err <- nil
}

type Request struct {
	Target string
	Action string
	Err    chan error
}

func main() {
}
