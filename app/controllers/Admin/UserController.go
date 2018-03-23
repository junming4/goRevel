package Admin

import (
	"github.com/revel/revel"
	"fmt"
)

type UserController struct {
	BaseController
}

func (c UserController) Index() revel.Result  {
	roleList := models.RoleList(0, "", "", "0,10", "")
	fmt.Println(roleList)
	c.Render(roleList)
	return c.RenderTemplate("backend/role/index.html")
}