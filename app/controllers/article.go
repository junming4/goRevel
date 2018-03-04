package controllers

import (
	"github.com/revel/revel"
	"fmt"
	"github.com/revel/revel/cache"
	"time"
	"myapp/app/models"
)

type Article struct {
	*revel.Controller
}

type Product struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

func (c Article) ArticleList() revel.Result {
	var product *Product
	name := revel.Config.StringDefault("cache.hosts", "localhost")


	fmt.Println(name)
	producesx := Product{Id:1,Name:"商品ID",Desc:"不是好东西"}

	cache.Set("jey",producesx,10)
	fmt.Println(product)

	produces := Product{Id:1,Name:"商品ID",Desc:"不是好东西"}
	cache.Add("jey", produces, 1)

	cache.Get("jey",product)
	fmt.Println(product)

	return c.RenderHTML("kkk")
}

func (c Article) TestRedis() revel.Result {
	redis := cache.NewRedisCache("127.0.0.1:6379","123456",10)

	err := redis.Set("int",1, time.Duration(-1))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var i int

	redis.Get("int", i)

	fmt.Println(i)

	return c.RenderHTML("kkk")
}

func (c Article) Show() revel.Result  {
	user := new(models.Users)
	rows,_ := db.Where("user_id=1").Rows()

	for rows.Next() {
		err := rows.Scan(user)
		if err != nil {
			fmt.Println(err)
			fmt.Println("kkk")
			return nil
		}
		fmt.Println(user.UserName)
	}


	db.Close()
	return c.RenderHTML("kkk")
}