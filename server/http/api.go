package http

import (
	"github.com/cilloparch/cillop/middlewares/i18n"
	"github.com/cilloparch/cillop/result"
	"github.com/gofiber/fiber/v2"
	"github.com/turistikrota/service.category/app/command"
	"github.com/turistikrota/service.category/app/query"
	"github.com/turistikrota/service.shared/server/http/auth/current_user"
)

func (h srv) CategoryCreate(ctx *fiber.Ctx) error {
	cmd := command.CategoryCreateCmd{}
	h.parseBody(ctx, &cmd)
	cmd.AdminUUID = current_user.Parse(ctx).UUID
	res, err := h.app.Commands.CategoryCreate(ctx.UserContext(), cmd)
	if err != nil {
		return result.Error(h.i18n.TranslateFromError(*err))
	}
	return result.SuccessDetail(Messages.Success.CategoryCreated, res)
}

func (h srv) CategoryUpdate(ctx *fiber.Ctx) error {
	detail := command.CategoryDetailCmd{}
	h.parseParams(ctx, &detail)
	cmd := command.CategoryUpdateCmd{}
	cmd.CategoryUUID = detail.CategoryUUID
	h.parseBody(ctx, &cmd)
	cmd.AdminUUID = current_user.Parse(ctx).UUID
	res, err := h.app.Commands.CategoryUpdate(ctx.UserContext(), cmd)
	if err != nil {
		return result.Error(h.i18n.TranslateFromError(*err))
	}
	return result.SuccessDetail(Messages.Success.CategoryUpdated, res)
}

func (h srv) CategoryUpdateOrder(ctx *fiber.Ctx) error {
	cmd := command.CategoryUpdateOrderCmd{}
	h.parseParams(ctx, &cmd)
	h.parseBody(ctx, &cmd)
	cmd.AdminUUID = current_user.Parse(ctx).UUID
	res, err := h.app.Commands.CategoryUpdateOrder(ctx.UserContext(), cmd)
	if err != nil {
		return result.Error(h.i18n.TranslateFromError(*err))
	}
	return result.SuccessDetail(Messages.Success.CategoryOrderUpdated, res)
}

func (h srv) CategoryEnable(ctx *fiber.Ctx) error {
	cmd := command.CategoryEnableCmd{}
	h.parseParams(ctx, &cmd)
	cmd.AdminUUID = current_user.Parse(ctx).UUID
	res, err := h.app.Commands.CategoryEnable(ctx.UserContext(), cmd)
	if err != nil {
		return result.Error(h.i18n.TranslateFromError(*err))
	}
	return result.SuccessDetail(Messages.Success.CategoryEnabled, res)
}

func (h srv) CategoryDisable(ctx *fiber.Ctx) error {
	cmd := command.CategoryDisableCmd{}
	h.parseParams(ctx, &cmd)
	cmd.AdminUUID = current_user.Parse(ctx).UUID
	res, err := h.app.Commands.CategoryDisable(ctx.UserContext(), cmd)
	if err != nil {
		return result.Error(h.i18n.TranslateFromError(*err))
	}
	return result.SuccessDetail(Messages.Success.CategoryDisabled, res)
}

func (h srv) CategoryDelete(ctx *fiber.Ctx) error {
	cmd := command.CategoryDeleteCmd{}
	h.parseParams(ctx, &cmd)
	cmd.AdminUUID = current_user.Parse(ctx).UUID
	res, err := h.app.Commands.CategoryDelete(ctx.UserContext(), cmd)
	if err != nil {
		return result.Error(h.i18n.TranslateFromError(*err))
	}
	return result.SuccessDetail(Messages.Success.CategoryDeleted, res)
}

func (h srv) CategoryAdminView(ctx *fiber.Ctx) error {
	query := query.CategoryFindQuery{}
	h.parseParams(ctx, &query)
	res, err := h.app.Queries.CategoryFind(ctx.UserContext(), query)
	if err != nil {
		return result.Error(h.i18n.TranslateFromError(*err))
	}
	return result.SuccessDetail(Messages.Success.CategoryAdminView, res)
}

func (h srv) CategoryView(ctx *fiber.Ctx) error {
	query := query.CategoryFindByUUIDQuery{}
	h.parseParams(ctx, &query)
	l, _ := i18n.GetLanguagesInContext(h.i18n, ctx)
	query.Locale = l
	res, err := h.app.Queries.CategoryFindByUUID(ctx.UserContext(), query)
	if err != nil {
		return result.Error(h.i18n.TranslateFromError(*err))
	}
	return result.SuccessDetail(Messages.Success.CategoryView, res)
}

func (h srv) CategoryFindFields(ctx *fiber.Ctx) error {
	q := query.CategoryFindFieldsByUUIDsQuery{}
	h.parseQuery(ctx, &q)
	res, err := h.app.Queries.CategoryFindFields(ctx.UserContext(), q)
	if err != nil {
		l, a := i18n.GetLanguagesInContext(h.i18n, ctx)
		return result.Error(h.i18n.TranslateFromError(*err, l, a))
	}
	return result.SuccessDetail(Messages.Success.CategoryView, res)
}

func (h srv) CategoryList(ctx *fiber.Ctx) error {
	q := query.CategoryFindAllQuery{}
	h.parseQuery(ctx, &q)
	res, err := h.app.Queries.CategoryFindAll(ctx.UserContext(), q)
	if err != nil {
		return result.Error(h.i18n.TranslateFromError(*err))
	}
	return result.SuccessDetail(Messages.Success.CategoryList, res.List)
}

func (h srv) CategoryListChild(ctx *fiber.Ctx) error {
	query := query.CategoryFindChildQuery{}
	h.parseParams(ctx, &query)
	res, err := h.app.Queries.CategoryFindChild(ctx.UserContext(), query)
	if err != nil {
		return result.Error(h.i18n.TranslateFromError(*err))
	}
	return result.SuccessDetail(Messages.Success.CategoryListChild, res.List)
}

func (h srv) CategoryAdminList(ctx *fiber.Ctx) error {
	q := query.CategoryAdminFindAllQuery{}
	h.parseQuery(ctx, &q)
	res, err := h.app.Queries.CategoryAdminFindAll(ctx.UserContext(), q)
	if err != nil {
		return result.Error(h.i18n.TranslateFromError(*err))
	}
	return result.SuccessDetail(Messages.Success.CategoryList, res)
}

func (h srv) CategoryAdminListChild(ctx *fiber.Ctx) error {
	query := query.CategoryFindChildQuery{}
	h.parseParams(ctx, &query)
	res, err := h.app.Queries.CategoryFindChild(ctx.UserContext(), query)
	if err != nil {
		return result.Error(h.i18n.TranslateFromError(*err))
	}
	return result.SuccessDetail(Messages.Success.CategoryListChild, res)
}

func (h srv) CategoryAdminListParents(ctx *fiber.Ctx) error {
	query := query.CategoryAdminFindParentsQuery{}
	h.parseQuery(ctx, &query)
	res, err := h.app.Queries.CategoryAdminFindParents(ctx.UserContext(), query)
	if err != nil {
		return result.Error(h.i18n.TranslateFromError(*err))
	}
	return result.SuccessDetail(Messages.Success.CategoryListChild, res)
}
