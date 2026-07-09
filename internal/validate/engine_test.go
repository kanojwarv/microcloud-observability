package validate

import (
	"context"
	"testing"
)

type fakeValidator struct {
	name string
}

func (f fakeValidator) Name() string {
	return f.name
}

func (f fakeValidator) Validate(context.Context) Result {
	return Result{
		Name:     f.name,
		Passed:   true,
		Score:    10,
		MaxScore: 10,
	}
}

func TestEngineRun(t *testing.T) {
	engine := New()

	engine.Register(fakeValidator{name: "one"})
	engine.Register(fakeValidator{name: "two"})

	report := engine.Run(context.Background())

	if len(report.Results) != 2 {
		t.Fatalf("expected 2 validators, got %d", len(report.Results))
	}

	if !report.Passed {
		t.Fatal("expected report to pass")
	}

	if report.Score != 20 {
		t.Fatalf("expected score 20, got %d", report.Score)
	}
}
