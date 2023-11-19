package event_stream

import (
	"github.com/cilloparch/cillop/events"
	"github.com/cilloparch/cillop/server"
	"github.com/turistikrota/service.category/app"
	"github.com/turistikrota/service.category/config"
)

type srv struct {
	app    app.Application
	topics config.Topics
	engine events.Engine
}

type Config struct {
	App    app.Application
	Engine events.Engine
	Topics config.Topics
}

func New(config Config) server.Server {
	return srv{
		app:    config.App,
		engine: config.Engine,
		topics: config.Topics,
	}
}

func (s srv) Listen() error {
	_ = s.engine.Subscribe(s.topics.Listing.Created, s.OnListingUpdated)
	_ = s.engine.Subscribe(s.topics.Listing.Updated, s.OnListingUpdated)
	return nil
}
