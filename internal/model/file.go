package model

type ModuleFile struct {
	Name           string
	RequireParsing bool

	SeparateSavePath bool
	// SavePath address where the file is saved
	SavePath []string
}
