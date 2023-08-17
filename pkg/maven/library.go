package maven

import "fmt"

type Library struct {
	GroupId    string `json:"group_id"`
	ArtefactId string `json:"artefact_id"`
}

func (l Library) String() string {
	return fmt.Sprintf("%s:%s", l.GroupId, l.ArtefactId)
}
