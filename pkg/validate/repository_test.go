package validate

import "testing"

func TestRepositoryValidator(t *testing.T) {
	_ = RepositoryValidator{Root: "../../"}.Validate()
}
