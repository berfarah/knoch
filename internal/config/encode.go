package config

type GeneralSettings struct {
	Workers int `toml:"parallel_workers"`
}

type EncodableConfig struct {
	General  GeneralSettings `toml:"general"`
	Projects []Project       `toml:"project"`
}

func (c *Config) Decode(ec *EncodableConfig) {
	c.Workers = ec.General.Workers
	for _, p := range ec.Projects {
		c.Projects.Add(p)
	}
}

func (c *Config) Encode() EncodableConfig {
	ec := EncodableConfig{
		General: GeneralSettings{
			Workers: c.Workers,
		},
		Projects: make([]Project, 0, len(c.Projects)),
	}

	for _, p := range c.Projects {
		ec.Projects = append(ec.Projects, p)
	}

	return ec
}
