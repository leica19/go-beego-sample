package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type ListController struct {
	beego.Controller
}

type CreateController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *ListController) Get() {
	c.TplName = "list.tpl"
}

func (c *CreateController) Get() {
	c.TplName = "create.tpl"
}
