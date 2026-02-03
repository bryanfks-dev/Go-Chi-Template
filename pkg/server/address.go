package server

func (srv *Server) Address() string {
	return "http://" + srv.server.Addr
}
