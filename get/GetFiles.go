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
	"strconv"

	. "github.com/gogufo/gufo-api-gateway/gufodao"
	"github.com/microcosm-cc/bluemonday"

	"github.com/getsentry/sentry-go"
	pb "github.com/gogufo/gufo-api-gateway/proto/go"
	"github.com/spf13/viper"
)

func GetFiles(t *pb.Request) (response *pb.Response) {

	ans := make(map[string]interface{})
	args := ToMapStringInterface(t.Args)
	p := bluemonday.UGCPolicy()

	if args["ownerid"] == nil {
		return ErrorReturn(t, 400, "000003", "Missing OwnerID")
	}

	ownerid := p.Sanitize(fmt.Sprintf("%s", args["ownerid"]))

	db, err := ConnectDBv2()
	if err != nil {
		if viper.GetBool("server.sentry") {
			sentry.CaptureException(err)
		} else {
			SetErrorLog(err.Error())
		}

		return ErrorReturn(t, 500, "000027", err.Error())
	}

	offset := 0
	limit := 25

	if args["offset"] != nil {
		offset, _ = strconv.Atoi(fmt.Sprintf("%v", args["offset"]))
	}

	if args["limit"] != nil {
		limit, _ = strconv.Atoi(fmt.Sprintf("%v", args["limit"]))
	}

	data := model.Files{}

	var count int64
	if args["ownerid"] == nil {
		db.Conn.Debug().Model(data).Where(`ownerid = ?`, ownerid).Count(&count)
		db.Conn.Debug().Where(`ownerid = ?`, ownerid).Limit(limit).Offset(offset).Find(&data)

	} else {
		group := p.Sanitize(fmt.Sprintf("%s", args["group"]))
		db.Conn.Debug().Model(data).Where(`ownerid = ? AND group = ?`, ownerid, group).Count(&count)
		db.Conn.Debug().Where(`ownerid = ? AND group = ?`, ownerid, group).Limit(limit).Offset(offset).Find(&data)
	}

	ans["files"] = data
	ans["filescount"] = count

	return response
}
