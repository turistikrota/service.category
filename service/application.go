package service

import (
	"github.com/cilloparch/cillop/events"
	"github.com/cilloparch/cillop/helpers/cache"
	"github.com/cilloparch/cillop/validation"
	"github.com/turistikrota/service.category/app"
	"github.com/turistikrota/service.category/app/command"
	"github.com/turistikrota/service.category/app/query"
	"github.com/turistikrota/service.category/config"
	"github.com/turistikrota/service.category/domains/category"
	"github.com/turistikrota/service.shared/db/mongo"
)

type Config struct {
	App         config.App
	EventEngine events.Engine
	Validator   *validation.Validator
	MongoDB     *mongo.DB
	CacheSrv    cache.Service
}

func NewApplication(cnf Config) app.Application {
	categoryFactory := category.NewFactory()
	categoryRepo := category.NewRepo(cnf.MongoDB.GetCollection(cnf.App.DB.Category.Collection), categoryFactory)
	categoryEvents := category.NewEvents(category.EventConfig{
		Publisher: cnf.EventEngine,
		Topics:    cnf.App.Topics,
	})
	categoryValidators := category.NewValidators(categoryFactory.Errors, cnf.App)

	return app.Application{
		Commands: app.Commands{
			CategoryCreate:       command.NewCategoryCreateHandler(categoryFactory, categoryRepo, categoryEvents),
			CategoryUpdate:       command.NewCategoryUpdateHandler(categoryFactory, categoryRepo, categoryEvents),
			CategoryDelete:       command.NewCategoryDeleteHandler(categoryRepo, categoryEvents),
			CategoryDisable:      command.NewCategoryDisableHandler(categoryRepo, categoryEvents),
			CategoryEnable:       command.NewCategoryEnableHandler(categoryRepo, categoryEvents),
			CategoryUpdateOrder:  command.NewCategoryUpdateOrderHandler(categoryRepo, categoryEvents),
			CategoryValidatePost: command.NewCategoryValidatePostHandler(categoryFactory, categoryRepo, categoryValidators, categoryEvents),
		},
		Queries: app.Queries{
			CategoryFind:           query.NewCategoryFindHandler(categoryRepo, cnf.CacheSrv, cnf.App),
			CategoryFindBySlug:     query.NewCategoryFindBySlugHandler(categoryRepo, cnf.CacheSrv, cnf.App),
			CategoryFindChild:      query.NewCategoryFindChildHandler(categoryRepo, cnf.CacheSrv),
			CategoryFindAll:        query.NewCategoryFindAllHandler(categoryRepo, cnf.CacheSrv),
			CategoryAdminFindChild: query.NewCategoryAdminFindChildHandler(categoryRepo, cnf.CacheSrv),
			CategoryAdminFindAll:   query.NewCategoryAdminFindAllHandler(categoryRepo, cnf.CacheSrv),
		},
	}
}
