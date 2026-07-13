package loader

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/kanojwarv/microcloud-observability/internal/model"
)

func LoadDashboards(
	dir string,
) ([]model.Dashboard, error) {

	var dashboards []model.Dashboard

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

		dashboard.Dashboard.Source = filepath.Join(
			dir,
			file.Name(),
		)

		dashboards = append(
			dashboards,
			dashboard.Dashboard,
		)
	}

	return dashboards, nil
}
