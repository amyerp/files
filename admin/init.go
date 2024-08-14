package admin

import (
gt "files/admin/get"
pt "files/admin/post"
	. "github.com/gogufo/gufo-api-gateway/gufodao"
	pb "github.com/gogufo/gufo-api-gateway/proto/go"
)

func Init(t *pb.Request) (response *pb.Response) {

switch *t.Method {
case "GET":
	response = gt.Init(t) // No such param
case "POST":
	response = pt.Init(t)
default:
	response = ErrorReturn(t, 404, "000012", "Missing argument")
}

return response

}
