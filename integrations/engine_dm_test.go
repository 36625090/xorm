// Copyright 2021 The Xorm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build dm
// +build dm

package integrations

import "github.com/36625090/xorm/schemas"

func init() {
	dbtypes = append(dbtypes, schemas.DAMENG)
}
