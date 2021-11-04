package network

type Config struct {
	DistributionLength int
	HiddenLength       int
	OutputLength       int
	Alpha              float64
}

const (
	DistributionLength = 36
	HiddenLength       = 10
	OutputLength       = 5
	Alpha              = 0.5
)

func NewPerceptronDefaultConfig() *Config {
	return &Config{
		DistributionLength: DistributionLength,
		HiddenLength:       HiddenLength,
		OutputLength:       OutputLength,
		Alpha:              Alpha,
	}
}
