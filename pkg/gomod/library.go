package gomod

type Library struct {
	Module string `json:"module"`
}

func (l Library) String() string {
	return l.Module
}
