// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto

import (
	"reflect"
	"testing"

	"github.com/youtube/vitess/go/bson"
	myproto "github.com/youtube/vitess/go/vt/mysqlctl/proto"
)

type reflectBinlogTransaction struct {
	Statements []reflectStatement
	Timestamp  int64
	GTIDField  myproto.GTIDField
}

type extraBinlogTransaction struct {
	Extra      int
	Statements []reflectStatement
	Timestamp  int64
	GTIDField  myproto.GTIDField
}

type reflectStatement struct {
	Category int
	Sql      []byte
}

func TestBinlogTransaction(t *testing.T) {
	reflected, err := bson.Marshal(&reflectBinlogTransaction{
		Statements: []reflectStatement{
			{
				Category: 1,
				Sql:      []byte("sql"),
			},
		},
		Timestamp: 456,
	})
	if err != nil {
		t.Error(err)
	}
	want := string(reflected)

	custom := BinlogTransaction{
		Statements: []Statement{
			{
				Category: 1,
				Sql:      []byte("sql"),
			},
		},
		Timestamp: 456,
	}
	encoded, err := bson.Marshal(&custom)
	if err != nil {
		t.Error(err)
	}
	got := string(encoded)
	if want != got {
		t.Errorf("want\n%#v, got\n%#v", want, got)
	}

	var unmarshalled BinlogTransaction
	err = bson.Unmarshal(encoded, &unmarshalled)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(custom, unmarshalled) {
		t.Errorf("%#v != %#v", custom, unmarshalled)
	}

	extra, err := bson.Marshal(&extraBinlogTransaction{})
	if err != nil {
		t.Error(err)
	}
	err = bson.Unmarshal(extra, &unmarshalled)
	if err != nil {
		t.Error(err)
	}
}

func TestStatementString(t *testing.T) {
	table := map[string]Statement{
		`{BL_UNRECOGNIZED: "SQL"}`: Statement{Category: BL_UNRECOGNIZED, Sql: []byte("SQL")},
		`{BL_BEGIN: "SQL"}`:        Statement{Category: BL_BEGIN, Sql: []byte("SQL")},
		`{BL_COMMIT: "SQL"}`:       Statement{Category: BL_COMMIT, Sql: []byte("SQL")},
		`{BL_ROLLBACK: "SQL"}`:     Statement{Category: BL_ROLLBACK, Sql: []byte("SQL")},
		`{BL_DML: "SQL"}`:          Statement{Category: BL_DML, Sql: []byte("SQL")},
		`{BL_DDL: "SQL"}`:          Statement{Category: BL_DDL, Sql: []byte("SQL")},
		`{BL_SET: "SQL"}`:          Statement{Category: BL_SET, Sql: []byte("SQL")},
		`{7: "SQL"}`:               Statement{Category: 7, Sql: []byte("SQL")},
	}
	for want, input := range table {
		if got := input.String(); got != want {
			t.Errorf("%#v.String() = %#v, want %#v", input, got, want)
		}
	}
}
