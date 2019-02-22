package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["ClusterManager/controllers:ClusterController"] = append(beego.GlobalControllerRouter["ClusterManager/controllers:ClusterController"],
        beego.ControllerComments{
            Method: "Role",
            Router: `/role`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ClusterManager/controllers:ClusterController"] = append(beego.GlobalControllerRouter["ClusterManager/controllers:ClusterController"],
        beego.ControllerComments{
            Method: "Status",
            Router: `/status`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ClusterManager/controllers:ClusterController"] = append(beego.GlobalControllerRouter["ClusterManager/controllers:ClusterController"],
        beego.ControllerComments{
            Method: "T",
            Router: `/test`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
