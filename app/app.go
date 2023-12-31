package app

import (
	"github.com/turistikrota/service.category/app/command"
	"github.com/turistikrota/service.category/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CategoryCreate          command.CategoryCreateHandler
	CategoryUpdate          command.CategoryUpdateHandler
	CategoryDelete          command.CategoryDeleteHandler
	CategoryDisable         command.CategoryDisableHandler
	CategoryEnable          command.CategoryEnableHandler
	CategoryUpdateOrder     command.CategoryUpdateOrderHandler
	CategoryValidateListing command.CategoryValidateListingHandler
}

type Queries struct {
	CategoryFindFields       query.CategoryFindFieldsByUUIDsHandler
	CategoryFind             query.CategoryFindHandler
	CategoryFindByUUID       query.CategoryFindByUUIDHandler
	CategoryFindChild        query.CategoryFindChildHandler
	CategoryFindAll          query.CategoryFindAllHandler
	CategoryAdminFindChild   query.CategoryAdminFindChildHandler
	CategoryAdminFindAll     query.CategoryAdminFindAllHandler
	CategoryAdminFindParents query.CategoryAdminFindParentsHandler
}
