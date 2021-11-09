package network

type Config struct {
	DistributionLength int
	HiddenLength       int
	OutputLength       int
	Alpha              float64
	Epsilon            float64
	WeightsStored      string
}

const (
	DistributionLength = 36
	HiddenLength       = 10
	OutputLength       = 5
	Alpha              = 0.05
	Epsilon            = 0.1
	WeightsStored      = "../../../resources/data/weights.data"
)

func NewRBFDefaultConfig() *Config {
	return &Config{
		DistributionLength: DistributionLength,
		HiddenLength:       HiddenLength,
		OutputLength:       OutputLength,
		Alpha:              Alpha,
		Epsilon:            Epsilon,
		WeightsStored:      WeightsStored,
	}
}
