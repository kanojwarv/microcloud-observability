package loader

import (
	"os"

	"go.yaml.in/yaml/v3"

	"github.com/kanojwarv/microcloud-observability/internal/model"
)

func LoadRecordingRuleFile(
	path string,
) (model.RecordingRuleFile, error) {

	var rules model.RecordingRuleFile

	data, err := os.ReadFile(path)
	if err != nil {
		return rules, err
	}

	err = yaml.Unmarshal(data, &rules)
	if err != nil {
		return rules, err
	}

	return rules, nil
}
