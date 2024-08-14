//////////////////////////////////////////////////////////////////////////////////
// Copyright 2021-2024 Alexey Yanchenko <mail@yanchenko.me>                          //
//                                                                              //
// This file is part of the ERP library.                                        //
//                                                                              //
//  Unauthorized copying of this file, via any media is strictly prohibited     //
//  Proprietary and confidential                                                //
//////////////////////////////////////////////////////////////////////////////////

package model

import (
	"gorm.io/gorm"
)

type Files struct {
	gorm.Model
	UUID        string `gorm:"column:uuid;type:varchar(60);UNIQUE;NOT NULL;"  json:"uuid"` //xwfdl
	OwnerID     string `gorm:"column:ownerid;type:varchar(60);NOT NULL;"  json:"ownerid"`  //rwefd
	Owner       string `gorm:"column:owner;type:varchar(254);NOT NULL;" json:"owner"`      //company
	Type        string `gorm:"column:type;type:varchar(254);NOT NULL;" json:"type"`        // avatar, file etc.
	Name        string `gorm:"column:name;type:varchar(254);NOT NULL;" json:"name"`        //br
	Ext         string `gorm:"column:ext;type:varchar(254);DEFAULT '';" json:"ext"`        //png
	Link        string `gorm:"column:link;type:varchar(254);DEFAULT '';" json:"link"`      //br-2020.png
	Sync        bool   `gorm:"column:sync;type:bool;DEFAULT '0';" json:"sync"`
	Description string `gorm:"column:description;type:varchar(254);DEFAULT '';" json:"description"` //business registration document
	Group       string `gorm:"column:group;type:longtext;DEFAULT '';" json:"group"`
	Comment     string `gorm:"column:comment;type:longtext;DEFAULT '';" json:"comment"`
}
