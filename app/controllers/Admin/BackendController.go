package Admin

import "github.com/revel/revel"

type BackendController struct {
	BaseController
}

func (c BackendController) Index() revel.Result {
	return c.RenderTemplate("backend/dashboard.html")
}
