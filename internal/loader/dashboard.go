package loader

import (
	"os"

	"go.yaml.in/yaml/v3"

	"github.com/kanojwarv/microcloud-observability/internal/model"
)

func LoadDashboardFile(
	path string,
) (model.DashboardFile, error) {

	var dashboard model.DashboardFile

	data, err := os.ReadFile(path)
	if err != nil {
		return dashboard, err
	}

	err = yaml.Unmarshal(data, &dashboard)
	if err != nil {
		return dashboard, err
	}

	return dashboard, nil
}
