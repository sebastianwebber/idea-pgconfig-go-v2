package config

const (
	// SBuffersMax contains the max value for the shared_buffers parameter
	SBuffersMax = 8 * GB

	// SBuffersMin contains the minimum value for the shared_buffers parameter
	SBuffersMin = 128 * KB
)

func sharedBuffers(i *Input) Parameter {

	out := Parameter{
		input: i,
		Name:  "shared_buffers",
		Type:  Bytes,
	}

	var paramValue uint64 = i.TotalRAM / 4

	if paramValue > SBuffersMax {
		paramValue = SBuffersMax
	}

	if paramValue < SBuffersMin {
		paramValue = SBuffersMin
	}

	out.value = paramValue

	return out
}
