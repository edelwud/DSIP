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
	Centers            = 4
	DeltaRate          = 0.1
	LastChangeRate     = 0.05
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
