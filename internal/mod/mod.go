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

		logger.StartSpinner("\tgo get "+module.Package, "✅\tGot "+module.Package)
		downloadModule(module.Package, app.ShortName, logger)
		logger.StopSpinner()
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
