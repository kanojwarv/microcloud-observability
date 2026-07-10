package loader

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/kanojwarv/microcloud-observability/internal/model"
)

func LoadDashboards(
	dir string,
) ([]model.DashboardFile, error) {

	var dashboards []model.DashboardFile

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {

		if file.IsDir() {
			continue
		}

		if !strings.HasSuffix(
			file.Name(),
			".dashboard.yaml",
		) {
			continue
		}

		dashboard, err := LoadDashboardFile(
			filepath.Join(
				dir,
				file.Name(),
			),
		)

		if err != nil {
			return nil, err
		}

		dashboards = append(
			dashboards,
			dashboard,
		)
	}

	return dashboards, nil
}
