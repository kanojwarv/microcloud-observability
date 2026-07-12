package loader

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadRecordingRuleFile(
	t *testing.T,
) {

	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf(
			"failed to get cwd: %v",
			err,
		)
	}

	path := filepath.Join(
		cwd,
		"..",
		"..",
		"recordingrules",
		"compute.rules.yaml",
	)

	rules, err := LoadRecordingRuleFile(
		path,
	)

	if err != nil {
		t.Fatalf(
			"unexpected error: %v",
			err,
		)
	}

	if len(rules.Groups) == 0 {
		t.Fatal(
			"expected groups",
		)
	}

	group := rules.Groups[0]

	if group.Name != "microcloud-compute" {
		t.Fatalf(
			"expected group %q, got %q",
			"microcloud-compute",
			group.Name,
		)
	}

	if len(group.Rules) == 0 {
		t.Fatal(
			"expected recording rules",
		)
	}
}

func TestLoadRecordingRuleFileMissing(
	t *testing.T,
) {

	_, err := LoadRecordingRuleFile(
		"does-not-exist.yaml",
	)

	if err == nil {
		t.Fatal(
			"expected error",
		)
	}
}
