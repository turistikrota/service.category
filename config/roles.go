package config

import "github.com/turistikrota/service.shared/base_roles"

type categoryRoles struct {
	Create       string
	Update       string
	Enable       string
	Disable      string
	Delete       string
	ViewAdmin    string
	List         string
	ListChildren string
	ReOrder      string
}

type roles struct {
	base_roles.Roles
	Category categoryRoles
}

var Roles = roles{
	Roles: base_roles.BaseRoles,
	Category: categoryRoles{
		Create:       "category.create",
		Update:       "category.update",
		Enable:       "category.enable",
		Disable:      "category.disable",
		Delete:       "category.delete",
		ViewAdmin:    "category.view.admin",
		List:         "category.list",
		ListChildren: "category.list.children",
		ReOrder:      "category.reorder",
	},
}
