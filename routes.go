package main

// Register routes on the server's router.
func (s *server) routes() {
	s.router.GET("/incoming-buses/:busStopCode/:serviceNumber", s.handleIncomingBusesGet)
}
