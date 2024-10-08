//////////////////////////////////////////////////////////////////////////////////
// Copyright 2024 Alexey Yanchenko <mail@yanchenko.me>                          //
//                                                                              //
// This file is part of the ERP library.                                        //
//                                                                              //
//  Unauthorized copying of this file, via any media is strictly prohibited     //
//  Proprietary and confidential                                                //
//////////////////////////////////////////////////////////////////////////////////

package get

//api/v3/files/getfile/{file_id}

//api/v3/files/getfiles?group={group}&ownerid={ownerid}
import (
	"files/model"
	"fmt"

	"github.com/getsentry/sentry-go"
	. "github.com/gogufo/gufo-api-gateway/gufodao"
	pb "github.com/gogufo/gufo-api-gateway/proto/go"
	"github.com/microcosm-cc/bluemonday"
	"github.com/spf13/viper"
)

func GetFile(t *pb.Request) (response *pb.Response) {

	ans := make(map[string]interface{})
	args := ToMapStringInterface(t.Args)

	db, err := ConnectDBv2()
	if err != nil {
		if viper.GetBool("server.sentry") {
			sentry.CaptureException(err)
		} else {
			SetErrorLog(err.Error())
		}

		return ErrorReturn(t, 500, "000027", err.Error())
	}

	p := bluemonday.UGCPolicy()

	if args["fileid"] == nil {
		return ErrorReturn(t, 400, "000003", "Missing FileID")
	}

	fileid := p.Sanitize(fmt.Sprintf("%s", args["fileid"]))

	file := model.Files{}
	db.Conn.Debug().Model(file).Where("uuid = ?", fileid).First(&file)

	ans["file"] = file
	response = Interfacetoresponse(t, ans)

	return response
}
