package controllers

import (
//	"database/sql"

	"github.com/revel/revel"
//	"github.com/jackc/pgx/v4"
)

type APIInfo struct {
    Name string ` json:"api" `
    Version string ` json:"version" `
	Stream string ` json:"stream" `
}

var defaultInfo = APIInfo{Name: "Reflow", Version: "v0.1", Stream: "DEV"}

type TechnicAPIController struct {
    *revel.Controller
    MyMappedData map[string]interface{}
}

func (c TechnicAPIController) ApiRoot() revel.Result {
	return c.RenderJSON(defaultInfo)
}