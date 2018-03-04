package Admin

import "github.com/revel/revel"

type PermissionController struct {
	BaseController
}

//列表
func (c PermissionController) PermissionIndex() revel.Result {
	return c.RenderTemplate("backend/login.html")
}

//创建
func (c PermissionController) PermissionCreate() revel.Result {
	return c.RenderTemplate("backend/login.html")
}

//保存
func PermissionStore(c *revel.Controller) revel.Result {
	return c.RenderTemplate("backend/login.html")
}

//修改
func PermissionEdit(c *revel.Controller) revel.Result {
	return c.RenderTemplate("backend/login.html")
}
//更新
func PermissionUpdate(c *revel.Controller) revel.Result {
	return c.RenderTemplate("backend/login.html")
}

//删除
func PermissionDestroy(c *revel.Controller) revel.Result {
	return c.RenderTemplate("backend/login.html")
}