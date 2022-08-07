package model

import (
	"strings"
)

type SurveyResult struct {
	Modules []string

	Name       string
	ShortName  string
	Version    string
	Logger     string
	DB         string
	HTTP       string
	Dockerfile bool
}

func EmptySurveyResult() *SurveyResult {
	return &SurveyResult{
		Logger:     None,
		DB:         None,
		HTTP:       None,
		Dockerfile: false,
	}
}

func (a *SurveyResult) Params() map[string]any {
	return map[string]any{
		"Name":      a.Name,
		"ShortName": a.ShortName,
		"Version":   a.Version,
		"HasLogger": a.Logger != None,
		"HasDB":     a.DB != None,
		"HasHTTP":   a.HTTP != None,
	}
}

func (a *SurveyResult) ExtractShortName() string {
	if strings.Contains(a.Name, "/") {
		split := strings.Split(a.Name, "/")
		return split[len(split)-1]
	} else {
		return a.Name
	}
}

// Execute extract some fields from existing fields
func (a *SurveyResult) Execute() *App {
	a.ShortName = a.ExtractShortName()

	return &App{
		Name:            a.Name,
		ShortName:       a.ShortName,
		Params:          a.Params(),
		SelectedModules: nil,
		RequiredModules: nil,
	}
}
