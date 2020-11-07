package options

type CorsOption struct {
	Origin        []string `yaml:"origin"`
	Method        []string `yaml:"method"`
	AllowHeader   []string `yaml:"allow_header"`
	ExposedHeader []string `yaml:"exposed_header"`
	MaxAge        int      `yaml:"max_age"`
	Credentials   bool     `yaml:"credentials"`
}