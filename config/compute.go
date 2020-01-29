package config

// Compute does the magic
func (i *Input) Compute() *[]Parameter {
	var output []Parameter

	if i.PostgresVersion == 0 {
		i.PostgresVersion = PGVersion
	}

	output = append(output,
		sharedBuffers(i))

	return &output
}
