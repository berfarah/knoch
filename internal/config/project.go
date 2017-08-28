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

func (set Projects) MarshalYAML() (interface{}, error) {
	projects := make([]Project, 0, len(set))
	for _, p := range set {
		projects = append(projects, p)
	}
	return projects, nil
}

func (set *Projects) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var projects []Project
	if err := unmarshal(&projects); err != nil {
		return err
	}

	for _, p := range projects {
		set.Add(p)
	}

	return nil
}
