/*
 * Copyright 2022 The Go Authors<36625090@qq.com>. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package xorm

// Config 数据库连接配置
type Config struct {
	Driver  string   `json:"driver" hcl:"driver"`
	Master  string   `json:"master" hcl:"master"` //root:123456@tcp(127.0.0.1:3306)/test?charset=utf8
	Slaves  []string `json:"slaves" hcl:"slaves"`
	ShowSql bool     `json:"show_sql" hcl:"show_sql"`
	MaxIdle int      `json:"max_idle" hcl:"max_idle"`
	MaxConn int      `json:"max_conn" hcl:"max_conn"`
}
