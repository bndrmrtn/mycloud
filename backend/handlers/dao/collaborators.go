package dao

type Collaborator struct {
	Email      string `json:"email"`
	Permission struct {
		Create bool `json:"create"`
		Read   bool `json:"read"`
		Update bool `json:"update"`
		Delete bool `json:"delete"`
	} `json:"permission"`
}

func (c *Collaborator) DoRemove() bool {
	return !(c.Permission.Delete || c.Permission.Update || c.Permission.Read || c.Permission.Create)
}
