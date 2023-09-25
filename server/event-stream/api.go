package event_stream

import (
	"context"
	"encoding/json"

	"github.com/turistikrota/service.category/app/command"
	"github.com/turistikrota/service.category/domains/post"
)

func (s srv) OnPostUpdated(data []byte) {
	e := post.PostUpdatedEvent{}
	err := json.Unmarshal(data, &e)
	if err != nil {
		return
	}
	s.app.Commands.CategoryValidatePost(context.Background(), command.CategoryValidatePostCmd{
		Post: e.Entity,
		User: e.User,
	})
}
