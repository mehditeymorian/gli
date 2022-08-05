package builder

import (
	"context"
	"fmt"
	"github.com/google/go-github/v45/github"
	"github.com/mehditeymorian/gli/internal/config"
	"github.com/mehditeymorian/gli/internal/model"
	"golang.org/x/oauth2"
	"html/template"
	"io"
	"log"
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
}

func NewBuilder(cfg config.Config) *Builder {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	return &Builder{
		Client: client,
		Config: cfg,
	}
}

func (b *Builder) Build(app *model.App) {
	b.ParentDirectory = app.ShortName

	b.DownloadModule(model.HTTP, app.HTTP)
	b.DownloadModule(model.DB, app.DB)
	b.DownloadModule(model.Logger, app.Logger)
	b.DownloadSingle("Dockerfile", app.Dockerfile, "", false, app)
	b.DownloadSingle("go.mod", true, "", true, app)
}

func (b *Builder) DownloadModule(module, technology string) {
	files := b.Config.ModuleTechnologyFiles(module, technology)

	if technology != "none" && files != nil {
		log.Printf("Module %s/%s\n", module, technology)

		dir := filepath.Join(b.ParentDirectory, "internal", module)

		for _, file := range files {
			log.Printf("Downloading %s\n", file)

			path := "template/" + module + "/" + technology + "/" + file
			reader, _, err := b.Client.Repositories.DownloadContents(context.Background(), user, repo, path, nil)
			handleErr("failed to download content from templates", err)

			if _, err := os.Stat(dir); os.IsNotExist(err) {
				err := os.MkdirAll(dir, os.ModePerm)
				handleErr("failed to create module directory", err)
			}

			file, err := os.Create(dir + "/" + file)
			handleErr("failed to create file", err)

			_, err = io.Copy(file, reader)
			handleErr("failed to copy content to file", err)

		}

	}
}

func (b *Builder) DownloadSingle(fileName string, required bool, directory string, requireParsing bool, app *model.App) {
	if !required {
		return
	}

	log.Printf("Downloading %s\n", fileName)

	path := "template/" + fileName

	if directory != "" {
		err := os.MkdirAll(directory, os.ModePerm)
		handleErr("failed to create directory", err)
	}

	var filePath string
	if directory != "" {
		filePath = filepath.Join(b.ParentDirectory, directory, fileName)
	} else {
		filePath = filepath.Join(b.ParentDirectory, fileName)
	}

	reader, _, err := b.Client.Repositories.DownloadContents(context.Background(), user, repo, path, nil)
	handleErr("failed to download content from templates", err)

	file, err := os.Create(filePath)
	handleErr("failed to create file", err)

	if requireParsing {
		buf := new(strings.Builder)
		_, err := io.Copy(buf, reader)
		handleErr("failed to copy content to buffer", err)

		temp := template.Must(template.New(fileName).Parse(buf.String()))

		params := app.Params()

		result := new(strings.Builder)

		err = temp.Execute(result, params)
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
