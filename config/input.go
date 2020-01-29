package config

const (
	B  = 1
	KB = B * 1024
	MB = KB * 1024
	GB = MB * 1024
)

// Input is foo
type Input struct {
	OS              string
	Arch            string
	TotalRAM        uint64
	Profile         string
	DiskType        string
	MaxConnections  int
	PostgresVersion float32
}
