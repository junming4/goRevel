package Admin

import "github.com/revel/revel"

type ExampleController struct {
	BaseController
}

func (c ExampleController) Index() revel.Result {
	return c.RenderTemplate("backend/_layouts/header.html")
}
