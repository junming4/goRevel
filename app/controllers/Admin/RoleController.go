package Admin

import (
	"github.com/revel/revel"
	"myapp/app/models"
	"fmt"
)

type RoleController struct {
	BaseController
}


//列表
func (c RoleController) Index() revel.Result {
	roleList := models.RoleList(0, "", "", "0,10", "")
	fmt.Println(roleList)
	c.Render(roleList)
	return c.RenderTemplate("backend/role/index.html")
}

//新增
func (c RoleController) Create() revel.Result {
	return  c.RenderTemplate("backend/role/create.html")
}

//保存数据
func (c RoleController) Store() revel.Result  {
	roleName := c.Params.Form.Get("roleName")
	roleSlug := c.Params.Form.Get("roleSlug")
	roleDescription := c.Params.Form.Get("roleDescription")

	c.Validation.Required(roleName).Message("角色名不能为空!")
	c.Validation.Required(roleSlug).Message("角色权限控制不能为空!")

	roles := models.Roles{0, roleName,
	roleSlug, roleDescription, 0, 11, "2018-03-11 12:39:42"}
	fmt.Println(models.CreateRole(roles))

	if models.CreateRole(roles) != true {
		return c.Redirect("/admin/role/index")
	}
	return c.Redirect("/admin/role/index")
}

//修改数据
func (c RoleController) Edit(id int) revel.Result {

	c.Validation.Min(id,1).Message("角色ID为空!")

	role := models.GetRole(id)

	c.Render(role)
	return c.RenderTemplate("backend/role/create.html")
}
