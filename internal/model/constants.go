package model

const (
	None       = "none"
	DB         = "db"
	HTTP       = "http"
	Logger     = "logger"
	Dockerfile = "Dockerfile"
)

type Flag string

const Verbose Flag = "verbose"
