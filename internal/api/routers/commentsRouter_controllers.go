package routers

import (
	beego "github.com/beego/beego/v2/server/web"

	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	beego.GlobalControllerRouter["eos_bot/api/props_crud/controllers:PropertiesController"] = append(beego.GlobalControllerRouter["eos_bot/api/props_crud/controllers:PropertiesController"],
	beego.ControllerComments{
	Method: "Post",
	Router: "/",
	AllowHTTPMethods: []string{"post"},
	MethodParams: param.Make(),
	Filters: nil,
	Params: nil})

	beego.GlobalControllerRouter["eos_bot/api/props_crud/controllers:PropertiesController"] = append(beego.GlobalControllerRouter["eos_bot/api/props_crud/controllers:PropertiesController"],
	beego.ControllerComments{
	Method: "GetAll",
	Router: "/",
	AllowHTTPMethods: []string{"get"},
	MethodParams: param.Make(),
	Filters: nil,
	Params: nil})

	beego.GlobalControllerRouter["eos_bot/api/props_crud/controllers:PropertiesController"] = append(beego.GlobalControllerRouter["eos_bot/api/props_crud/controllers:PropertiesController"],
	beego.ControllerComments{
	Method: "GetOne",
	Router: "/:id",
	AllowHTTPMethods: []string{"get"},
	MethodParams: param.Make(),
	Filters: nil,
	Params: nil})

	beego.GlobalControllerRouter["eos_bot/api/props_crud/controllers:PropertiesController"] = append(beego.GlobalControllerRouter["eos_bot/api/props_crud/controllers:PropertiesController"],
	beego.ControllerComments{
	Method: "Put",
	Router: "/:id",
	AllowHTTPMethods: []string{"put"},
	MethodParams: param.Make(),
	Filters: nil,
	Params: nil})

	beego.GlobalControllerRouter["eos_bot/api/props_crud/controllers:PropertiesController"] = append(beego.GlobalControllerRouter["eos_bot/api/props_crud/controllers:PropertiesController"],
	beego.ControllerComments{
	Method: "Delete",
	Router: "/:id",
	AllowHTTPMethods: []string{"delete"},
	MethodParams: param.Make(),
	Filters: nil,
	Params: nil})
	
	beego.GlobalControllerRouter["eos_bot/api/props_crud/controllers:PropertiesController"] = append(beego.GlobalControllerRouter["eos_bot/api/props_crud/controllers:PropertiesController"],
	beego.ControllerComments{
	Method: "GetAnal",
	Router: "/analysis",
	AllowHTTPMethods: []string{"get"},
	MethodParams: param.Make(),
	Filters: nil,
	Params: nil})

}
