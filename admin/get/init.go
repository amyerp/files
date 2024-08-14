package get

import (
	. "github.com/gogufo/gufo-api-gateway/gufodao"
	pb "github.com/gogufo/gufo-api-gateway/proto/go"
)

func Init(t *pb.Request) (response *pb.Response) {

	switch *t.ParamID {
	case "cronstatus":
		response = CheckCron(t)
	default:
		response = ErrorReturn(t, 404, "000012", "Missing argument")
	}

	return response
}
