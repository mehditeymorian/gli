package mod

import (
	"github.com/mehditeymorian/gli/internal/logger"
	"github.com/mehditeymorian/gli/internal/model"
	"os/exec"
)

func DownloadModules(app *model.App, logger logger.Logger) {
	for _, module := range app.SelectedModules {
		if module.Package == "" {
			continue
		}

		logger.Printf("Downloading Module %s:\n", module.Name)
		logger.PrintfV("go get %s\n", module.Package)
		downloadModule(module.Package, app.ShortName, logger)
	}

}

func downloadModule(module, projectDirectory string, logger logger.Logger) {
	cmd := exec.Command("go", "get", module)
	cmd.Dir = projectDirectory + "/"

	output, err := cmd.CombinedOutput()
	if err != nil {
		logger.Printf("an error occurred during downloading module %s: %s\n", module, err.Error())
	}

	logger.PrintfV("%s\n", output)
}
