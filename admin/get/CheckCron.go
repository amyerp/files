package get

import (
	"fmt"
	. "files/global"

	. "github.com/gogufo/gufo-api-gateway/gufodao"
	pb "github.com/gogufo/gufo-api-gateway/proto/go"
	"github.com/spf13/viper"
)

func CheckCron(t *pb.Request) (response *pb.Response) {

	ans := make(map[string]interface{})
	setingskey := fmt.Sprintf("%s.cron", MicroServiceName)
	isCron := viper.GetBool(setingskey)
	ans["cronstatus"] = fmt.Sprintf("%t", isCron)

	return Interfacetoresponse(t, ans)
}
