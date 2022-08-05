package mod

import (
	"github.com/mehditeymorian/gli/internal/config"
	"github.com/mehditeymorian/gli/internal/model"
	"log"
	"os/exec"
)

func DownloadModules(cfg config.Config, app *model.App) {
	requiredModules := app.RequiredModules(cfg.Modules)

	for _, module := range requiredModules {
		log.Printf("Downloading Module %s:\n", module)
		downloadModule(module, app.ShortName)
	}

}

func downloadModule(module, projectDirectory string) {
	cmd := exec.Command("go", "get", module)
	cmd.Dir = projectDirectory + "/"

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("an error occured during downloading module %s: %s\n", module, err.Error())
	}

	log.Printf("%s\n", output)
}
