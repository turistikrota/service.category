package post

type PostUpdatedEvent struct {
	Entity *Entity    `json:"entity"`
	User   UserDetail `json:"user"`
}
