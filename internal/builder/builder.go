package builder

import (
	"context"
	"fmt"
	"github.com/google/go-github/v45/github"
	"github.com/mehditeymorian/gli/internal/config"
	"github.com/mehditeymorian/gli/internal/logger"
	"github.com/mehditeymorian/gli/internal/model"
	"golang.org/x/oauth2"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const (
	user  = "mehditeymorian"
	repo  = "gli"
	token = "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"
)

type Builder struct {
	Client          *github.Client
	Config          config.Config
	ParentDirectory string
	Params          map[string]any
	Logger          logger.Logger
}

func NewBuilder(cfg config.Config, logger logger.Logger) *Builder {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	return &Builder{
		Client: client,
		Config: cfg,
		Logger: logger,
	}
}

func (b *Builder) Build(app *model.App) {
	b.ParentDirectory = app.ShortName
	b.Params = app.Params

	for _, module := range app.SelectedModules {
		b.DownloadModule(module)
	}

	for _, module := range app.RequiredModules {
		b.DownloadModule(module)
	}
}

func (b *Builder) DownloadModule(module model.Module) {
	files := module.Files

	moduleName := "Module " + module.Name
	b.Logger.StartSpinner("\tDownloading "+moduleName, "✅\t"+moduleName+" Downloaded")

	savePath := module.GetSavePath(b.ParentDirectory)

	for _, file := range files {
		b.Logger.PrintfV("Downloading %s\n", file.Name)

		fileDownloadURL := module.DownloadURL + file.Name
		reader, _, err := b.Client.Repositories.DownloadContents(context.Background(), user, repo, fileDownloadURL, nil)
		handleErr("failed to download content from templates", err)

		b.SaveFile(reader, savePath, file)
	}

	b.Logger.StopSpinner()
}

func (b *Builder) SaveFile(reader io.ReadCloser, savePath string, moduleFile model.ModuleFile) {
	if _, err := os.Stat(savePath); os.IsNotExist(err) {
		err := os.MkdirAll(savePath, os.ModePerm)
		handleErr("failed to create savePath", err)
	}

	filePath := filepath.Join(savePath, moduleFile.Name)

	file, err := os.Create(filePath)
	handleErr("failed to create file", err)

	if moduleFile.RequireParsing {
		buf := new(strings.Builder)
		_, err := io.Copy(buf, reader)
		handleErr("failed to copy content to buffer", err)

		temp := template.Must(template.New(moduleFile.Name).Parse(buf.String()))

		result := new(strings.Builder)

		err = temp.Execute(result, b.Params)
		handleErr("failed to execute template", err)

		_, err = file.Write([]byte(result.String()))
	} else {
		_, err = io.Copy(file, reader)
		handleErr("failed to copy content to file", err)
	}
}

func handleErr(message string, err error) {
	if err != nil {
		panic(fmt.Errorf("%s: %w", message, err))
	}
}
