package mod

import (
	"github.com/mehditeymorian/gli/internal/model"
	"log"
	"os/exec"
)

func DownloadModules(app *model.App) {
	for _, module := range app.SelectedModules {
		if module.Package == "" {
			continue
		}

		log.Printf("Downloading Module %s:\n", module.Name)
		log.Printf("go get %s\n", module.Package)
		downloadModule(module.Package, app.ShortName)
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
