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

type DownloadStatus int

const (
	FullyDownloaded DownloadStatus = iota
	PartiallyDownloaded
	NothingDownloaded
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

func (b *Builder) Build(app *model.App) (int, int, int) {
	b.ParentDirectory = app.ShortName
	b.Params = app.Params

	b.Logger.Title("Downloading Template")

	totalDownloaded := 0
	partiallyDownloaded := 0

	modules := append(app.SelectedModules, app.RequiredModules...)

	for _, module := range modules {
		status := b.DownloadModule(module)
		if status == FullyDownloaded {
			totalDownloaded++
		} else if status == PartiallyDownloaded {
			partiallyDownloaded++
		}
	}

	return totalDownloaded, partiallyDownloaded, len(app.SelectedModules) + len(app.RequiredModules)
}

func (b *Builder) DownloadModule(module model.Module) DownloadStatus {
	files := module.Files

	moduleName := "Module " + module.Name
	b.Logger.StartSpinner("\tDownloading " + moduleName)

	savePath := module.GetSavePath(b.ParentDirectory)

	downloaded := 0

	for _, file := range files {
		b.Logger.SetSpinnerMessage("-> " + file.Name)
		b.Logger.PrintfV("Downloading %s\n", file.Name)

		fileDownloadURL := module.DownloadURL + file.Name
		reader, _, err := b.Client.Repositories.DownloadContents(context.Background(), user, repo, fileDownloadURL, nil)
		if err != nil {
			b.Logger.PrintfV("failed to download content from templates: %s\n", err.Error())

			continue
		}

		err = b.SaveFile(reader, savePath, file)
		if err != nil {
			b.Logger.PrintfV("failed to save file %s%s: %s\n", savePath, file.Name, err.Error())

			continue
		}

		downloaded++
	}

	var status DownloadStatus
	var msg string

	if downloaded == 0 {
		msg = "🤕\t" + moduleName + " Didn't Download"
		status = NothingDownloaded
	} else if downloaded < len(files) {
		msg = "😥\t" + moduleName + " Partially Downloaded"
		status = PartiallyDownloaded
	} else {
		// some case except downloaded == len(files) also match in case of error check here.
		msg = "✅\t" + moduleName + " Downloaded"
		status = FullyDownloaded
	}

	b.Logger.StopSpinner(msg)

	return status
}

func (b *Builder) SaveFile(reader io.ReadCloser, savePath string, moduleFile model.ModuleFile) error {
	if _, err := os.Stat(savePath); os.IsNotExist(err) {
		err := os.MkdirAll(savePath, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create savePath: %w", err)
		}
	}

	filePath := filepath.Join(savePath, moduleFile.Name)

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}

	if moduleFile.RequireParsing {
		buf := new(strings.Builder)
		_, err := io.Copy(buf, reader)
		if err != nil {
			return fmt.Errorf("failed to copy content to buffer: %w", err)
		}

		funcMap := template.FuncMap{
			"toUpper": strings.ToUpper,
		}

		temp := template.Must(template.New(moduleFile.Name).Funcs(funcMap).Parse(buf.String()))

		result := new(strings.Builder)

		err = temp.Execute(result, b.Params)
		if err != nil {
			return fmt.Errorf("failed to execute template: %w", err)
		}

		_, err = file.Write([]byte(result.String()))
		if err != nil {
			return fmt.Errorf("failed to write to file: %w", err)
		}
	} else {
		_, err = io.Copy(file, reader)
		if err != nil {
			return fmt.Errorf("failed to copy content to file: %w", err)
		}
	}

	return nil
}
