package post

import (
	"fmt"
	"files/cron"
	. "files/global"

	. "github.com/gogufo/gufo-api-gateway/gufodao"
	pb "github.com/gogufo/gufo-api-gateway/proto/go"
	"github.com/microcosm-cc/bluemonday"
	"github.com/spf13/viper"
)

func UpdateCron(t *pb.Request) (response *pb.Response) {
	ans := make(map[string]interface{})
	args := ToMapStringInterface(t.Args)
	p := bluemonday.UGCPolicy()

	if args["action"] == nil {
		fmt.Printf("Missing important data")
		return ErrorReturn(t, 404, "000012", "Missing important data")
	}

	action := p.Sanitize(fmt.Sprintf("%v", args["action"]))
	setingskey := fmt.Sprintf("%s.cron", MicroServiceName)

	if action == "true" {
		viper.Set(setingskey, true)
		/// Run Cron
		go cron.Init()
	} else {
		viper.Set(setingskey, false)
	}

	ans["answer"] = action
	return Interfacetoresponse(t, ans)
}
