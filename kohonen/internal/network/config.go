package network

type Config struct {
	DistributionLength int
	OutputLength       int
	Centers            int
	Beta               float64
}

const (
	DistributionLength = 36
	OutputLength       = 5
	Centers            = 2
	Beta               = 0.1
)

func NewKohonenDefaultConfig() *Config {
	return &Config{
		DistributionLength: DistributionLength,
		OutputLength:       OutputLength,
		Centers:            Centers,
		Beta:               Beta,
	}
}
