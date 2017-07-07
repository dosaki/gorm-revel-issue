package controllers

import (
	"github.com/revel/revel"
	"gorm_issue/app/datasource"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
    animal := datasource.Animal{Age: 2, Name: "Gopher"}
    datasource.DB.Create(&animal)

	return c.Render()
}
