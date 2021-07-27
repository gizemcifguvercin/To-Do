package commands

type CreateToDoCommand struct {
	Id             string     `json:"id"`
	Type           string     `json:"type"`
	Name           string     `json:"name"`
}

