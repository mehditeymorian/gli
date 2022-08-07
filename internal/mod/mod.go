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
		if len(module.Package) == 0 {
			continue
		}

		total += len(module.Package)

		for _, modulePackage := range module.Package {
			logger.StartSpinner("\tgo get " + modulePackage)
			err := downloadModule(modulePackage, app.ShortName, logger)
			if err != nil {
				logger.PrintfV("an error occurred during downloading module %s: %s\n", module, err.Error())
				logger.StopSpinner("🤕\tFailed to Get " + modulePackage)
			} else {
				logger.StopSpinner("✅\tGot " + modulePackage)
				totalDownloaded++
			}
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
