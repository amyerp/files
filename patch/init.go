//////////////////////////////////////////////////////////////////////////////////
// Copyright 2024 Alexey Yanchenko <mail@yanchenko.me>                          //
//                                                                              //
// This file is part of the ERP library.                                        //
//                                                                              //
//  Unauthorized copying of this file, via any media is strictly prohibited     //
//  Proprietary and confidential                                                //
//////////////////////////////////////////////////////////////////////////////////

package patch

//api/v3/files/getfilebyid/{file_id}
//api/v3/files/getfilesbyowner/{ownerid}
//api/v3/files/getfilesbygroup/{group}
import (

	//	. "github.com/gogufo/gufo-api-gateway/gufodao"
	pb "github.com/gogufo/gufo-api-gateway/proto/go"
)

func Init(t *pb.Request) (response *pb.Response) {
	return response
}
