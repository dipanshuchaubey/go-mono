package types

type (
	Route struct {
		URL       string `yaml:"url"`
		Method    string `yaml:"method"`
		Handler   string `yaml:"handler"`
		Timeout   int    `yaml:"timeout"`
		Auth      bool   `yaml:"auth"`
		RateLimit int    `yaml:"rateLimit"`
	}

	Routes map[string]Route

	RoutesMap map[string]map[string]Route
)
