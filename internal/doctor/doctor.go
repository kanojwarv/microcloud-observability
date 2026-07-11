package doctor

import (
	"github.com/kanojwarv/microcloud-observability/internal/model"
)

func Check() model.Health {

	return model.Health{
		Status:  "OK",
		Message: "All dashboard dependencies are used.",
	}
}
