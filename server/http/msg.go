package http

type successMessages struct {
	CategoryCreated      string
	CategoryUpdated      string
	CategoryOrderUpdated string
	CategoryDeleted      string
	CategoryEnabled      string
	CategoryDisabled     string
	CategoryView         string
	CategoryAdminView    string
	CategoryList         string
	CategoryListChild    string
}

type errorMessages struct {
	RequiredAuth      string
	CurrentUserAccess string
	AdminRoute        string
}

type messages struct {
	Success successMessages
	Error   errorMessages
}

var Messages = messages{
	Success: successMessages{
		CategoryCreated:      "http_success_category_created",
		CategoryUpdated:      "http_success_category_updated",
		CategoryOrderUpdated: "http_success_category_order_updated",
		CategoryDeleted:      "http_success_category_deleted",
		CategoryEnabled:      "http_success_category_enabled",
		CategoryDisabled:     "http_success_category_disabled",
		CategoryView:         "http_success_category_view",
		CategoryAdminView:    "http_success_category_admin_view",
		CategoryList:         "http_success_category_list",
		CategoryListChild:    "http_success_category_list_child",
	},
	Error: errorMessages{
		RequiredAuth:      "http_error_required_auth",
		CurrentUserAccess: "http_error_current_user_access",
		AdminRoute:        "http_error_admin_route",
	},
}
