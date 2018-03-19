package Admin

import (
	"github.com/revel/revel"
	//"myapp/app/models"
	//"fmt"
	"myapp/app/models"
	//"fmt"
	"fmt"
)

type AuthController struct {
	*revel.Controller
}

func (c AuthController) Login() revel.Result {

	user := models.GetUserById(2)
	fmt.Println(user)

	users := models.Users{0, "22eee222333", "222233", "222", "333", 2}

	models.CreateUser(users)

	role := models.Roles{0, "3333", "kk", "", 1, 1111, ""}
	models.CreateRole(role)
	/*user := new(models.Users)
	Db.Find(&user)
	user := models.Users{UserName:"kkdkd"}

	db.NewRecord(user) // => 主键为空返回`true`
	db.Create(&user)

	fmt.Println(user)*/
	//fmt.Println(db.Create(&user))
	if _, ok := c.Session["UserId"]; ok {
		return c.Redirect(AuthController.Login)
	}
	return c.RenderTemplate("backend/login.html")
}

func (c AuthController) PostLogin() revel.Result {
	userName := c.Params.Form.Get("userName")
	password := c.Params.Form.Get("password")
	//os.Exit(0)
	c.Validation.Required(userName).Message("名称不能为空!")
	c.Validation.Required(password).Message("密码不能为空!")
	return c.Redirect("")
}
