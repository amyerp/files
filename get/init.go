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
	. "github.com/gogufo/gufo-api-gateway/gufodao"
	pb "github.com/gogufo/gufo-api-gateway/proto/go"
)

func Init(t *pb.Request) (response *pb.Response) {

	switch *t.Param {
	case "getfile":
		response = GetFile(t)
	case "getfiles":
		response = GetFiles(t)
	default:
		response = ErrorReturn(t, 404, "000010", "Missing  Param")
	}

	return response
}
