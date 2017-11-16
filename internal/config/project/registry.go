package project

import "sort"

type Registry map[string]Project

var Tracker = Registry{}

func (r Registry) Add(p Project) bool {
	_, found := r[p.Dir]
	r[p.Dir] = p
	return !found
}

func (r Registry) Remove(p Project) {
	delete(r, p.Dir)
}

func (r Registry) Sorted() []Project {
	var temp []string
	var sorted []Project

	for dir := range r {
		temp = append(temp, dir)
	}

	sort.Strings(temp)

	for _, key := range temp {
		sorted = append(sorted, r[key])
	}

	return sorted
}

func Register(p Project) bool {
	return Tracker.Add(p)
}

func Remove(p Project) bool {
	if _, ok := Tracker[p.Dir]; !ok {
		return false
	}

	Tracker.Remove(p)
	return true
}

func Fetch(dir string) (Project, bool) {
	project, ok := Tracker[dir]
	return project, ok
}
