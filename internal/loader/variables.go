package loader

import (
	"os"

	"go.yaml.in/yaml/v3"

	"github.com/kanojwarv/microcloud-observability/internal/model"
)

func LoadVariablesFile(
	path string,
) (model.VariablesFile, error) {

	var variables model.VariablesFile

	data, err := os.ReadFile(path)
	if err != nil {
		return variables, err
	}

	err = yaml.Unmarshal(data, &variables)
	if err != nil {
		return variables, err
	}

	return variables, nil
}
