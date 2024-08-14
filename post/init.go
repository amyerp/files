//////////////////////////////////////////////////////////////////////////////////
// Copyright 2024 Alexey Yanchenko <mail@yanchenko.me>                          //
//                                                                              //
// This file is part of the ERP library.                                        //
//                                                                              //
//  Unauthorized copying of this file, via any media is strictly prohibited     //
//  Proprietary and confidential                                                //
//////////////////////////////////////////////////////////////////////////////////

package post

import (
	"encoding/json"
	"files/model"

	"github.com/getsentry/sentry-go"
	. "github.com/gogufo/gufo-api-gateway/gufodao"
	pb "github.com/gogufo/gufo-api-gateway/proto/go"
	"github.com/spf13/viper"
)

// api/v3/files/file
func Init(t *pb.Request) (response *pb.Response) {

	switch *t.Param {
	case "file":
		response = AddFile(t)
	default:
		response = ErrorReturn(t, 404, "000010", "Missing  Param")
	}

	return response

}

func AddFile(t *pb.Request) (response *pb.Response) {
	ans := make(map[string]interface{})
	args := ToMapStringInterface(t.Args)
	//p := bluemonday.UGCPolicy()

	if args["ownerid"] == nil || args["name"] == nil || args["ext"] == nil {
		return ErrorReturn(t, 406, "000012", "Missing  Important data")
	}

	db, err := ConnectDBv2()
	if err != nil {
		if viper.GetBool("server.sentry") {
			sentry.CaptureException(err)
		} else {
			SetErrorLog(err.Error())
		}

		return ErrorReturn(t, 500, "000027", err.Error())
	}

	data := &model.Files{}

	JsonArgs, err := json.Marshal(args)
	if err != nil {
		return ErrorReturn(t, 500, "000028", err.Error())
	}

	err = json.Unmarshal(JsonArgs, &data)
	if err != nil {
		return ErrorReturn(t, 500, "000028", err.Error())
	}
	dataid := Hashgen(12)
	data.UUID = dataid

	err = db.Conn.Create(&data).Error
	if err != nil {
		return ErrorReturn(t, 400, "000005", err.Error())
	}

	ans["uuid"] = dataid
	response = Interfacetoresponse(t, ans)
	return response

}
