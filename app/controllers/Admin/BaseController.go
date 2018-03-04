package Admin

import "github.com/revel/revel"

// 公用Controller, 其它Controller继承它
type BaseController struct {
	*revel.Controller
}

func checkUser(c BaseController) revel.Result {
	if userId := c.GetUserId(); userId == "" {
		c.Flash.Error("请先登录")
		//return c.Redirect(Admin.Login.index)
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

//退出登录
func (c BaseController) Logout() {
	delete(c.Session, "userId")
	delete(c.Session, "Username")
}

func init() {
	//revel.InterceptFunc(checkUser, revel.BEFORE, &BaseController{})
}