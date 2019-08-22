package api

import context "context"

type Server struct {
}

func (s *Server) Ping(ctx context.Context, req *Void) (*CheckMessage, error) {
	return &CheckMessage{Check: "OK"}, nil
}
