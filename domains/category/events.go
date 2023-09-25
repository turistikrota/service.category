package category

import (
	"fmt"

	"github.com/cilloparch/cillop/events"
	"github.com/turistikrota/service.category/config"
	"github.com/turistikrota/service.category/domains/post"
)

type Events interface {
	Created(event CreatedEvent)
	Updated(event UpdatedEvent)
	Enabled(event EnabledEvent)
	Disabled(event DisabledEvent)
	Deleted(event DeletedEvent)
	UpdateOrder(event OrderUpdatedEvent)
	PostValidationSuccess(event PostValidationSuccessEvent)
	PostValidationFailed(event PostValidationFailedEvent)
}

type (
	CreatedEvent struct {
		AdminUUID string  `json:"admin_uuid"`
		Entity    *Entity `json:"entity"`
	}
	UpdatedEvent struct {
		AdminUUID string  `json:"admin_uuid"`
		Entity    *Entity `json:"entity"`
	}
	EnabledEvent struct {
		AdminUUID  string `json:"admin_uuid"`
		EntityUUID string `json:"entity_uuid"`
	}
	OrderUpdatedEvent struct {
		AdminUUID  string `json:"admin_uuid"`
		EntityUUID string `json:"entity_uuid"`
		Order      int16  `json:"order"`
	}
	DisabledEvent struct {
		AdminUUID  string `json:"admin_uuid"`
		EntityUUID string `json:"entity_uuid"`
	}
	DeletedEvent struct {
		AdminUUID  string `json:"admin_uuid"`
		EntityUUID string `json:"entity_uuid"`
	}
	PostValidationSuccessEvent struct {
		PostUUID string          `json:"postUUID"`
		Post     *post.Entity    `json:"entity"`
		User     UserDetailEvent `json:"user"`
	}
	PostValidationFailedEvent struct {
		PostUUID string                  `json:"postUUID"`
		Post     *post.Entity            `json:"entity"`
		Errors   []*post.ValidationError `json:"errors"`
		User     UserDetailEvent         `json:"user"`
	}
	UserDetailEvent struct {
		UUID string `json:"uuid"`
		Name string `json:"name"`
		Code string `json:"code"`
	}
)

type categoryEvents struct {
	publisher events.Publisher
	topics    config.Topics
}

type EventConfig struct {
	Topics    config.Topics
	Publisher events.Publisher
}

func NewEvents(cnf EventConfig) Events {
	return &categoryEvents{
		publisher: cnf.Publisher,
		topics:    cnf.Topics,
	}
}

func (e *categoryEvents) Created(event CreatedEvent) {
	_ = e.publisher.Publish(e.topics.Category.Created, &event)
}

func (e *categoryEvents) Updated(event UpdatedEvent) {
	_ = e.publisher.Publish(e.topics.Category.Updated, &event)
}

func (e *categoryEvents) Enabled(event EnabledEvent) {
	_ = e.publisher.Publish(e.topics.Category.Enabled, &event)
}

func (e *categoryEvents) Disabled(event DisabledEvent) {
	_ = e.publisher.Publish(e.topics.Category.Disabled, &event)
}

func (e *categoryEvents) Deleted(event DeletedEvent) {
	_ = e.publisher.Publish(e.topics.Category.Deleted, &event)
}

func (e *categoryEvents) UpdateOrder(event OrderUpdatedEvent) {
	_ = e.publisher.Publish(e.topics.Category.OrderUpdated, &event)
}

func (e *categoryEvents) PostValidationSuccess(event PostValidationSuccessEvent) {
	fmt.Println(fmt.Sprintf("PostValidationSuccessEvent event fired to %s", e.topics.Category.PostValidationSuccess))
	_ = e.publisher.Publish(e.topics.Category.PostValidationSuccess, &event)
}

func (e *categoryEvents) PostValidationFailed(event PostValidationFailedEvent) {
	fmt.Println(fmt.Sprintf("PostValidationFailedEvent event fired to %s", e.topics.Category.PostValidationFailed))
	_ = e.publisher.Publish(e.topics.Category.PostValidationFailed, &event)
}
