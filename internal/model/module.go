package model

import "path"

type Module struct {
	// Name of the module
	Name string

	// DownloadURL location of module in the template to download from
	DownloadURL string

	// SavePath path that files are stored
	SavePath []string

	// Package module package to download
	Package string

	Files []ModuleFile
}

func (m Module) GetSavePath(parentDirectory string) string {
	join := make([]string, len(m.SavePath)+1)

	join[0] = parentDirectory

	for i, each := range m.SavePath {
		join[i+1] = each
	}

	return path.Join(join...)
}
