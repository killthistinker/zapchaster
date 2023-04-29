package server

func Init() {
	serManager := Configure()
	InitEndpoints(serManager)
}
