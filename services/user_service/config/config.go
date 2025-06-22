package config

type (
	Service struct {
		Name    string `yaml:"name"`
		Port    string `yaml:"port"`
		Timeout int    `yaml:"timeout"`
	}

	DB struct {
		URI         string `yaml:"uri"`
		Credentails string `yaml:"credentials_path"`
	}

	Database struct {
		Driver      string `yaml:"driver"`
		Master      DB     `yaml:"master"`
		ReadReplica DB     `yaml:"read_replica"`
	}

	Config struct {
		Service  Service  `yaml:"Service"`
		Database Database `yaml:"Database"`
	}
)
