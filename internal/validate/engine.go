package validate

import "context"

type Engine struct {
	validators []Validator
}

func New() *Engine {
	return &Engine{}
}

func (e *Engine) Register(v Validator) {
	e.validators = append(e.validators, v)
}

func (e *Engine) Run(ctx context.Context) Report {
	report := Report{
		Passed: true,
	}

	for _, validator := range e.validators {
		result := validator.Validate(ctx)

		report.Results = append(report.Results, result)

		report.Score += result.Score
		report.MaxScore += result.MaxScore

		if !result.Passed {
			report.Passed = false
		}
	}

	return report
}
