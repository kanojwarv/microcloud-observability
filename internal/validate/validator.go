package validate

import "context"

type Validator interface {
	Name() string
	Validate(context.Context) Result
}
