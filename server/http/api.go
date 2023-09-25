package http

import (
	"github.com/cilloparch/cillop/result"
	"github.com/gofiber/fiber/v2"
	"github.com/turistikrota/service.category/app/command"
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
	return nil
}

func (h srv) CategoryUpdateOrder(ctx *fiber.Ctx) error {
	return nil
}

func (h srv) CategoryEnable(ctx *fiber.Ctx) error {
	return nil
}

func (h srv) CategoryDisable(ctx *fiber.Ctx) error {
	return nil
}

func (h srv) CategoryDelete(ctx *fiber.Ctx) error {
	return nil
}

func (h srv) CategoryAdminView(ctx *fiber.Ctx) error {
	return nil
}

func (h srv) CategoryView(ctx *fiber.Ctx) error {
	return nil
}

func (h srv) CategoryList(ctx *fiber.Ctx) error {
	return nil
}

func (h srv) CategoryListChild(ctx *fiber.Ctx) error {
	return nil
}

func (h srv) CategoryAdminList(ctx *fiber.Ctx) error {
	return nil
}

func (h srv) CategoryAdminListChild(ctx *fiber.Ctx) error {
	return nil
}
