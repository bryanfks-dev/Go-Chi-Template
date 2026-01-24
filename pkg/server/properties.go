package server

func (srv *Server) Address() string {
	return srv.server.Addr
}
