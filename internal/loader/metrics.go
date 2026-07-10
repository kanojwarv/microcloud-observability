package loader

import (
	"os"

	"gopkg.in/yaml.v3"

	"github.com/kanojwarv/microcloud-observability/internal/model"
)

func LoadMetricsFile(path string) (model.MetricsFile, error) {
	var metrics model.MetricsFile

	data, err := os.ReadFile(path)
	if err != nil {
		return metrics, err
	}

	err = yaml.Unmarshal(data, &metrics)
	if err != nil {
		return metrics, err
	}

	return metrics, nil
}
