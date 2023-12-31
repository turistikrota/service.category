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
	CategoryFindFieldsByUUIDs query.CategoryFindFieldsByUUIDsHandler
	CategoryFindFieldsBySlugs query.CategoryFindFieldsBySlugsHandler
	CategoryFind              query.CategoryFindHandler
	CategoryFindBySlug        query.CategoryFindBySlugHandler
	CategoryFindChildByUUID   query.CategoryFindChildByUUIDHandler
	CategoryFindChildBySlug   query.CategoryFindChildBySlugHandler
	CategoryFindAllByUUIDs    query.CategoryFindAllByUUIDsHandler
	CategoryFindAllBySlugs    query.CategoryFindAllBySlugHandler
	CategoryAdminFindChild    query.CategoryAdminFindChildHandler
	CategoryAdminFindAll      query.CategoryAdminFindAllHandler
	CategoryAdminFindParents  query.CategoryAdminFindParentsHandler
}
