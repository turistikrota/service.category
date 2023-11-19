package listing

type ListingUpdatedEvent struct {
	Entity *Entity    `json:"entity"`
	User   UserDetail `json:"user"`
}
