package event_stream

import (
	"context"
	"encoding/json"

	"github.com/turistikrota/service.category/app/command"
	"github.com/turistikrota/service.category/domains/listing"
)

func (s srv) OnListingUpdated(data []byte) {
	e := listing.ListingUpdatedEvent{}
	err := json.Unmarshal(data, &e)
	if err != nil {
		return
	}
	s.app.Commands.CategoryValidateListing(context.Background(), command.CategoryValidateListingCmd{
		Listing: e.Entity,
		User:    e.User,
	})
}
