package loader

import (
	"os"
	"path/filepath"

	"go.yaml.in/yaml/v3"

	"github.com/kanojwarv/microcloud-observability/internal/model"
)

func LoadMetricsFile(
	path string,
) (model.MetricsFile, error) {

	var metricsFile model.MetricsFile

	data, err := os.ReadFile(
		path,
	)

	if err != nil {
		return metricsFile, err
	}

	err = yaml.Unmarshal(
		data,
		&metricsFile,
	)

	if err != nil {
		return metricsFile, err
	}

	return metricsFile, nil
}

func LoadMetrics(
	repositoryRoot string,
) ([]model.Metric, error) {

	var metrics []model.Metric

	files, err := filepath.Glob(
		filepath.Join(
			repositoryRoot,
			"metrics",
			"*.yaml",
		),
	)

	if err != nil {
		return nil, err
	}

	for _, file := range files {

		metricsFile, err := LoadMetricsFile(
			file,
		)

		if err != nil {
			return nil, err
		}

		for _, metric := range metricsFile.Metrics {

			metric.Source = file

			metrics = append(
				metrics,
				metric,
			)
		}
	}

	return metrics, nil
}
