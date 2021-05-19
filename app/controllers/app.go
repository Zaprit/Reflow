package controllers

import (
	gorpController "github.com/revel/modules/orm/gorp/app/controllers"
	"github.com/revel/revel"
)

type App struct {
	*gorpController.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}
