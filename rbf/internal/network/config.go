package network

type Config struct {
	DistributionLength int
	OutputLength       int
	Centers            int
	DeltaRate          float64
	LastChangeRate     float64
}

const (
	DistributionLength = 36
	OutputLength       = 5
	Centers            = 2
	DeltaRate          = 0.001
	LastChangeRate     = 0.0005
)

func NewRBFDefaultConfig() *Config {
	return &Config{
		DistributionLength: DistributionLength,
		OutputLength:       OutputLength,
		Centers:            Centers,
		DeltaRate:          DeltaRate,
		LastChangeRate:     LastChangeRate,
	}
}
