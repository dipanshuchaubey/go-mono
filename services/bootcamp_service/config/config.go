package config

type (
	Service struct {
		Name    string `yaml:"name"`
		Port    string `yaml:"port"`
		Timeout int    `yaml:"timeout"`
	}

	Credentials struct {
		BootcampAPI string `yaml:"bootcamp_api"`
	}

	EndPoints struct {
		PostLogin    string `yaml:"post_login"`
		GetBootcamps string `yaml:"get_bootcamps"`
		PostBootcamp string `yaml:"post_bootcamp"`
	}

	Config struct {
		Service     Service     `yaml:"Service"`
		Credentials Credentials `yaml:"Credentials"`
		EndPoints   EndPoints   `yaml:"EndPoints"`
	}
)
