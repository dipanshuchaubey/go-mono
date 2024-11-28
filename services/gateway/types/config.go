package types

type (
	Service struct {
		Name             string `yaml:"name"`
		Port             string `yaml:"port"`
		Timeout          int    `yaml:"timeout"`
		RoutesConfigPath string `yaml:"routes_config_path"`
	}

	EndPoints struct {
		GrpcUserService     string `yaml:"grpc_user_service"`
		GrpcBootcampService string `yaml:"grpc_bootcamp_service"`
	}

	Config struct {
		Service   Service   `yaml:"Service"`
		EndPoints EndPoints `yaml:"EndPoints"`
	}
)
