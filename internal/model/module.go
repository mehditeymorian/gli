package model

import "path"

type Module struct {
	// Name of the module
	Name string

	// DownloadURL location of module in the template to download from
	DownloadURL string

	// SavePath path that files are stored if ModuleFile#SeparateSavePath is false.
	SavePath []string

	// Package module package to download
	Package []string

	Files []ModuleFile
}

func (m Module) GetSavePath(parentDirectory string, file ModuleFile) string {
	join := []string{parentDirectory}

	savePath := m.SavePath
	if file.SeparateSavePath {
		savePath = file.SavePath
	}

	for _, each := range savePath {
		join = append(join, each)
	}

	return path.Join(join...)
}
