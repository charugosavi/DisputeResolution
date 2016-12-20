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
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strconv"
)

func (this *HDLS) load(db *shim.ChaincodeStub, key string) string {
	bytes, _ := db.GetState(key)
	return string(bytes)
}
func (this *HDLS) loadInt(db *shim.ChaincodeStub, key string) int {
	val, _ := strconv.Atoi(this.load(db, key))
	return val
}

func (this *HDLS) save(db *shim.ChaincodeStub, key string, val string) {
	db.PutState(key, []byte(val))
}
func (this *HDLS) saveInt(db *shim.ChaincodeStub, key string, val int) {
	this.save(db, key, strconv.Itoa(val))
}
func (this *HDLS) increase(db *shim.ChaincodeStub, key string, delta int) {
	this.saveInt(db, key, this.loadInt(db, key)+delta)
}
func (this *HDLS) decrease(db *shim.ChaincodeStub, key string, delta int) {
	this.increase(db, key, -delta)
}
func (this *HDLS) retInAdd(ok bool, err error) ([]byte, error) {
	if err != nil {
		return nil, fmt.Errorf("operation failed. %s", err)
	}
	if !ok {
		return nil, errors.New("operation failed. Row with given key already exists")
	}
	return nil, nil
}

func must(x interface{}, err error) interface{} {
	return x
}
func toInt32(x string) int32 {
	i, err := strconv.Atoi(x)
	if err != nil {
		return -1
	}
	return int32(i)
}
