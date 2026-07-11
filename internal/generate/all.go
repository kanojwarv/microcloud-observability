package generate

func All(
	repositoryRoot string,
) error {

	_, err := GraphJSON(
		repositoryRoot,
	)

	if err != nil {
		return err
	}

	_, err = GraphHTML(
		repositoryRoot,
	)

	if err != nil {
		return err
	}

	return nil
}
