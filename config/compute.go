package config

// Compute does the magic
func (i *Input) Compute() *Output {
	var output Output

	if i.PostgresVersion == 0 {
		i.PostgresVersion = PGVersion
	}

	memory := Category{
		Name: "Resource Usage / Memory",
	}

	sBuffers := sharedBuffers(i)
	sBuffers.setValue()

	mConn := maxConnections(i)
	mConn.setValue()
	memory.Parameters = append(memory.Parameters,
		sBuffers,
		mConn)

	output.Data = append(output.Data, memory)

	return &output
}

type Category struct {
	Name       string `json:"name"`
	Parameters []Parameter
}
