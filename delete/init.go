//////////////////////////////////////////////////////////////////////////////////
// Copyright 2024 Alexey Yanchenko <mail@yanchenko.me>                          //
//                                                                              //
// This file is part of the ERP library.                                        //
//                                                                              //
//  Unauthorized copying of this file, via any media is strictly prohibited     //
//  Proprietary and confidential                                                //
//////////////////////////////////////////////////////////////////////////////////

package delete

import (
	"files/model"
	"fmt"

	. "github.com/gogufo/gufo-api-gateway/gufodao"

	"github.com/getsentry/sentry-go"
	pb "github.com/gogufo/gufo-api-gateway/proto/go"
	"github.com/microcosm-cc/bluemonday"
	"github.com/spf13/viper"
)

func Init(t *pb.Request) (response *pb.Response) {
	ans := make(map[string]interface{})
	args := ToMapStringInterface(t.Args)

	p := bluemonday.UGCPolicy()

	if args["fileid"] == nil {
		return ErrorReturn(t, 400, "000003", "Missing FileID")
	}

	fileid := p.Sanitize(fmt.Sprintf("%s", args["fileid"]))

	db, err := ConnectDBv2()
	if err != nil {
		if viper.GetBool("server.sentry") {
			sentry.CaptureException(err)
		} else {
			SetErrorLog(err.Error())
		}

		return ErrorReturn(t, 500, "000027", err.Error())
	}

	db.Conn.Delete(model.Files{}, "uuid = ?", fileid)
	ans["response"] = "200501" // Business deleted
	response = Interfacetoresponse(t, ans)
	return response

}
