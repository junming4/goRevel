package Admin

import (
	"github.com/revel/revel"
	"myapp/app/models"
	//"fmt"
)

type AuthController struct {
	*revel.Controller
}

func (c AuthController) Login() revel.Result {
	if _, ok := c.Session["UserId"]; ok {
		return c.Redirect(BackendController.Index)
	}
	return c.RenderTemplate("backend/login.html")
}

func (c AuthController) PostLogin() revel.Result {
	userName := c.Params.Form.Get("userName")
	password := c.Params.Form.Get("password")
	c.Validation.MinSize(userName,1).Message("用户名不能为空!")
	c.Validation.MinSize(password,1).Message("密码不能为空!")

	//展示错误
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(AuthController.Login)
	}
	user := models.GetUserByUser(userName, password)
	if user == nil {
		c.Flash.Error("密码或者用户非法")
		return c.Redirect(AuthController.Login)
	}
	c.Session["UserId"] = string(user.UserId)
	c.Session["UserName"] = string(user.UserName)
	return c.Redirect(AuthController.Login)
}

//退出登录
func (c AuthController) Logout() revel.Result {
	delete(c.Session, "UserId")
	delete(c.Session, "UserName")
	return c.Redirect(AuthController.Login)
}
