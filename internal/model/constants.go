package model

const (
	None             = "none"
	DB               = "db"
	HTTP             = "http"
	Logger           = "logger"
	Dockerfile       = "Dockerfile"
	StartPoint       = "StartPoint"
	StartPointCli    = "Cli"
	StartPointSimple = "Simple"
)

type Flag string

const Verbose Flag = "verbose"
