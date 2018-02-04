package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {

	return c.Render()
}

func (c App) Blogs() revel.Result {
	return  c.RenderHTML("ok")
}

func (c App) Form() revel.Result {
	return  c.Render()
}

func (c App) CheckName() revel.Result {
	name := c.Params.Form.Get("name")
	//os.Exit(0)
	c.Validation.Required(name).Message("名称不能为空yyy!")
	return nil
}

//功能拦截
func CheckName2(c *revel.Controller) revel.Result {
	name := c.Params.Form.Get("name")
	//os.Exit(0)
	c.Validation.Required(name).Message("名称不能为空yyy!")
	return nil
}

type Person struct{
	Name string `json:"name"`
	Age int `json:"age"`
}

type PersonTwo struct {
	Name string `xml:"name"`
	Age int `xml:"age"`
}

//返回json 格式
func (c App) ReturnJson() revel.Result  {
	person := Person{Name:"小明",Age:21}
	return  c.RenderJSON(person)
}

func (c App) ReturnXml() revel.Result{
	person := PersonTwo{Name:"小明",Age:21}
	return  c.RenderXML(person)
}



func (c App) PostForm(request revel.Request) revel.Result {
	c.Flash.Error("非法输入!")
	name := c.Params.Form.Get("name")
	c.Validation.Required(name).Message("名称不能为空!")
	//c.Validation.MinSize(name, 3).Message("Your name is not long enough!")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Form)
	}

	person := Person{Name:name,Age:10}
	return c.Render(person)
}

func init() {
	revel.InterceptFunc(CheckName2,revel.BEFORE,App{}) //功能拦截
	revel.InterceptMethod(App.CheckName, revel.BEFORE) //方法拦截
}


