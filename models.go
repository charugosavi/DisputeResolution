/*
Licensed Materials - Property of IBM

6949-63F

(C) Copyright IBM Corp. 2016, 2016
All Rights Reserved

US Government Users Restricted Rights - Use, duplication or
disclosure restricted by GSA ADP Schedule Contract with IBM Corp.
*/

package main

import (
	"encoding/json"
	//"errors"
	"sort"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const (
	DEFAULT_ROWS_LIMIT = 10000
)

type Reference struct {
	Id  string
	Ids []string
}
type References struct {
	Data []Reference
}

// TODO: move to model_core.go
func (this *HDLS) putA(model string, id string, entity interface{}) error {
	entityBytes, err := json.Marshal(entity)
	if err != nil {
		return err
	}

	_, err = this.retInAdd(this.db.InsertRow(model, shim.Row{[]*shim.Column{
		&shim.Column{&shim.Column_String_{id}},
		&shim.Column{&shim.Column_String_{string(entityBytes)}},
	}}))

	return err
}

func (this *HDLS) delete(a, key string) error {
	err := this.db.DeleteRow(a, this.keyCol(key))
	return err
}

func (this *HDLS) val(row *shim.Row, x interface{}) error {
	jsonStr := row.Columns[1].GetString_()
	return json.Unmarshal([]byte(jsonStr), x)
}

//func (this *HDLS) listAllRows(model string) ([]*shim.Row, error) {
//	close, rowChannel, err := this.db.GetRows2(model, []shim.Column{})
//	if err != nil {
//		return nil, err
//	}
//
//	var rows []*shim.Row
//	for {
//		select {
//		case row, ok := <-rowChannel:
//			if !ok {
//				rowChannel = nil
//			} else {
//				rows = append(rows, &row)
//			}
//		}
//
//		// TODO make it available to controle the limit
//		if rowChannel == nil || len(rows) >= DEFAULT_ROWS_LIMIT {
//			break
//		}
//	}
//	close()
//
//	sort.Sort(ById(rows))
//	return rows, nil
//}

// TODO: above is the original by Ohsawa-san, which assumes the fix in shim.GetRows() as GetRows2()
// the following code is not fixed version.
// this is incorporated in order to avoid compilation error temporary.
func (this *HDLS) listAllRows(model string) ([]*shim.Row, error) {
	rowChannel, err := this.db.GetRows(model, []shim.Column{})
	if err != nil {
		return nil, err
	}

	var rows []*shim.Row
	for {
		select {
		case row, ok := <-rowChannel:
			if !ok {
				rowChannel = nil
			} else {
				rows = append(rows, &row)
			}
		}

		// TODO make it available to controle the limit
		if rowChannel == nil || len(rows) >= DEFAULT_ROWS_LIMIT {
			break
		}
	}
	// close()

	sort.Sort(ById(rows))
	return rows, nil
}

func (this *HDLS) keyCol(name string) []shim.Column {
	col1Val := name
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: col1Val}}
	columns = append(columns, col1)

	return columns
}

func (this *HDLS) getA(model, id string, entity interface{}) error {
	row, err := this.db.GetRow(model, this.keyCol(id))
	if err != nil {
		return err
	}
	if row.Columns == nil || len(row.Columns) < 2 {
		//return errors.New("NOT_FOUND|A requested entity with an ID " + id + " was not found")
		return nil
	}

	jsonStr := row.Columns[1].GetString_()

	//var entity interface{}
	err = json.Unmarshal([]byte(jsonStr), entity)
	if err != nil {
		return err
	}

	return nil
}

type ById []*shim.Row

func (s ById) Len() int {
	return len(s)
}
func (s ById) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ById) Less(i, j int) bool {
	iid := s[i].Columns[0].GetString_()
	jid := s[j].Columns[0].GetString_()
	return iid < jid
}

func remove(slice []string, search string) []string {
	result := []string{}
	for _, item := range slice {
		if item != search {
			result = append(result, item)
		}
	}
	return result
}
