package config

type Project struct {
	Repo string `yaml:"repo"`
	Dir  string `yaml:"dir"`
}

type Projects map[string]Project

func (set *Projects) Add(p Project) bool {
	_, found := (*set)[p.Repo]
	(*set)[p.Repo] = p
	return !found
}

func (set *Projects) Remove(p Project) {
	delete(*set, p.Repo)
}
