package config

import (
	"encoding/json"
)

type Project struct {
	Repo string `json:"repo"`
	Dir  string `json:"dir"`
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

func (set Projects) MarshalJSON() ([]byte, error) {
	projects := make([]Project, 0, len(set))
	for _, p := range set {
		projects = append(projects, p)
	}
	return json.Marshal(projects)
}

func (set *Projects) UnmarshalJSON(b []byte) error {
	var projects []Project
	if err := json.Unmarshal(b, &projects); err != nil {
		return err
	}

	for _, p := range projects {
		set.Add(p)
	}

	return nil
}
