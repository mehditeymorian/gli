package mod

import (
	"github.com/mehditeymorian/gli/internal/logger"
	"github.com/mehditeymorian/gli/internal/model"
	"os/exec"
)

func DownloadModules(app *model.App, logger logger.Logger) (int, int) {
	logger.Title("Downloading Packages")

	totalDownloaded := 0
	total := 0

	for _, module := range app.SelectedModules {
		if module.Package == "" {
			continue
		}

		total++

		logger.StartSpinner("\tgo get " + module.Package)
		err := downloadModule(module.Package, app.ShortName, logger)
		if err != nil {
			logger.PrintfV("an error occurred during downloading module %s: %s\n", module, err.Error())
			logger.StopSpinner("🤕\tFailed to Get " + module.Package)
		} else {
			logger.StopSpinner("✅\tGot " + module.Package)
			totalDownloaded++
		}
	}

	return totalDownloaded, total
}

func downloadModule(module, projectDirectory string, logger logger.Logger) error {
	cmd := exec.Command("go", "get", module)
	cmd.Dir = projectDirectory + "/"

	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	logger.PrintfV("%s\n", output)
	return nil
}
