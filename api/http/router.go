package http

func Init() {
	server, err := NewServer()
	if err != nil {
		server.l.Info(err)
	}
	server.Start()
}

func Route() {
	
}