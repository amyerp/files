// Copyright 2020 - 2024 Alexey Yanchenko <mail@yanchenko.me>
//
// This file is part of the Gufo library.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package entrypoint

import (
	. "files/model"

	. "github.com/gogufo/gufo-api-gateway/gufodao"
	"github.com/spf13/viper"
)

func CheckDBStructure() {
	//Check DB and table config
	db, err := ConnectDBv2()
	if err != nil {
		SetErrorLog("dbstructure.go:81: " + err.Error())
		//return "error with db"
	}

	dbtype := viper.GetString("database.type")

	if !db.Conn.Migrator().HasTable(&Files{}) {
		if dbtype == "mysql" {
			db.Conn.Set("gorm:table_options", "ENGINE=InnoDB;").Migrator().CreateTable(&Files{})
		} else {
			db.Conn.Migrator().CreateTable(&Files{})
		}
	}
}
