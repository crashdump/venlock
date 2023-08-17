package npm

type Library struct {
	Name string `json:"name"`
}

func (l Library) String() string {
	return l.Name
}
