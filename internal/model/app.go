package model

type App struct {
	Name            string
	ShortName       string
	Params          map[string]any
	SelectedModules []Module
	RequiredModules []Module
	Flags           map[Flag]bool
}
