package main

import (
	. "github.com/gogufo/gufo-api-gateway/gufodao"
	pb "github.com/gogufo/gufo-api-gateway/proto/go"

	ad "files/admin"
	dl "files/delete"
	gt "files/get"
	patch "files/patch"
	pt "files/post"
	. "files/version"
)

func Init(t *pb.Request) (response *pb.Response) {

	if t.UID == nil {
		response = ErrorReturn(t, 401, "000011", "You are not authorised")
		return response
	}

	switch *t.Param {
	case "admin":
		return admincheck(t)
	}

	if *t.Method == "GET" {

		switch *t.Param {
		case "info":
			response = info(t)
		case "health":
			response = health(t)
		default:
			response = ErrorReturn(t, 406, "000012", "Wrong request")
		}
	}

	switch *t.Method {
	case "GET":
		response = gt.Init(t)
	case "POST":
		response = pt.Init(t)
	case "PATCH":
		response = patch.Init(t)
	case "DELETE":
		response = dl.Init(t)
	default:
		response = ErrorReturn(t, 404, "00004", "Wrong Method")

	}

	return response

}

func info(t *pb.Request) (response *pb.Response) {
	ans := make(map[string]interface{})
	ans["pluginname"] = "files"
	ans["version"] = VERSIONPLUGIN
	ans["description"] = ""
	response = Interfacetoresponse(t, ans)
	return response
}

func health(t *pb.Request) (response *pb.Response) {
	ans := make(map[string]interface{})
	ans["health"] = "OK"
	response = Interfacetoresponse(t, ans)
	return response
}

func admincheck(t *pb.Request) (response *pb.Response) {

	if *t.IsAdmin != 1 {
		response = ErrorReturn(t, 401, "000012", "You have no admin rights")
	}

	return ad.Init(t)

}
