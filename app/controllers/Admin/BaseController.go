package Admin

import "github.com/revel/revel"

// 公用Controller, 其它Controller继承它
type BaseController struct {
	*revel.Controller
}

func (c BaseController) checkUser() revel.Result {
	revel.AppLog.Error("进入了这里")
	if userId := c.GetUserId(); userId == "" {
		c.Flash.Error("请先登录")
		return c.Redirect(AuthController.Login)
	}
	return nil
}

func (c BaseController) GetUserId() string {
	if userId, ok := c.Session["UserId"]; ok {
		return userId
	}
	return ""
}


func (c BaseController) HasAdmin() bool {
	return c.GetUserId() != ""
}


func init() {
	revel.InterceptMethod(BaseController.checkUser, revel.BEFORE)
}