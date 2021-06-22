package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"git.huawei.com/poissonsearch/wienerchain/contract/sdk"
	"git.huawei.com/poissonsearch/wienerchain/contract/sdk/logger"
	"git.huawei.com/poissonsearch/wienerchain/contract/sdk/smstub"
	"git.huawei.com/poissonsearch/wienerchain/proto/common"
)

type example01 struct {
}

var log = logger.GetDefaultLogger()

func (e example01) Init(stub sdk.ContractStub) common.InvocationResponse {
	fmt.Printf("Enter example01 init function\n")
	args := stub.Parameters()
	const numOfArgs = 2
	if len(args) < numOfArgs {
		return sdk.Error("Init parameter is not correct")
	}
	key := args[0]
	value := args[1]
	err := stub.PutKV(string(key), value)
	if err != nil {
		return sdk.Error("Init put kv failed")
	}

	return sdk.Success(nil)
}

func (e example01) Invoke(stub sdk.ContractStub) common.InvocationResponse {
	funcName := stub.FuncName()
	args := stub.Parameters()

	switch funcName {
	case "initMarble":
		return initMarble(stub, args)
	case "initMarbleCommon":
		return initMarbleCommon(stub, args)
	case "getMarbleCommon":
		return getMarbleCommon(stub, args)
	case "getMarble":
		return getMarble(stub, args)
	case "delMarble":
		return deleteMarble(stub, args)
	case "initRange":
		return initRange(stub, args)
	case "getRange":
		return getRange(stub, args)
	case "getMarbleCom":
		return getMarbleCom(stub, args)
	case "initMarbleAndIndex":
		return initMarbleAndIndex(stub, args)
	case "transferMarblesByIndex":
		return transferMarblesByIndex(stub, args)
	case "getPartRangeAndPutKV":
		return getPartRangeAndPutKV(stub, args)
	case "getJSONMarble":
		return getJSONMarble(stub)
	case "getRangeAndPutKV":
		return getRangeAndPutKV(stub, args)
	case "writeMarble":
		return writeMarble(stub, args)
	case "deleteComIndexOneRow":
		return deleteComIndexOneRow(stub, args)
	}
	str := fmt.Sprintf("Func name is not correct, the function name is %s ", funcName)
	return sdk.Error(str)
}

func deleteComIndexOneRow(stubInterface sdk.ContractStub, args [][]byte) common.InvocationResponse {
	const numOfArgs = 3
	if len(args) < numOfArgs {
		return sdk.Error("the number of args is not correct")
	}

	indexName := string(args[0])
	attr := string(args[1])
	objectName := string(args[2])

	err := stubInterface.DelComIndexOneRow(indexName, []string{attr}, objectName)
	if err != nil {
		return sdk.Error(err.Error())
	}
	return sdk.Success(nil)
}

func initMarbleCommon(stubInterface sdk.ContractStub, args [][]byte) common.InvocationResponse {
	const numOfArgs = 4
	if len(args) < numOfArgs {
		fmt.Printf("The args number for init marble is not correct\n")
		return sdk.Error("The args number for init marble is not correct")
	}
	marbleName := string(args[0])
	marbleOwner := string(args[1])
	marbleColor := string(args[2])
	marbleSizeStr := string(args[3])
	marbleSize, err := strconv.Atoi(marbleSizeStr)
	if err != nil {
		fmt.Printf("The marble size is not int type\n")
		return sdk.Error("The marble size is not int type")
	}

	marbleInfo := &marble{Name: marbleName, Color: marbleColor, Size: marbleSize, Owner: marbleOwner, ObjectType: "marble"}

	err = stubInterface.PutKVCommon(marbleInfo.Name, marbleInfo)
	if err != nil {
		return sdk.Error(err.Error())
	}

	return sdk.Success(nil)
}

func getMarbleCommon(stubInterface sdk.ContractStub, args [][]byte) common.InvocationResponse {
	if len(args) < 1 {
		fmt.Printf("enter getMarbleCommon\n")
		return sdk.Error("not enough parameters")
	}
	key := string(args[0])
	value, err := stubInterface.GetKV(key)
	if err != nil {
		fmt.Printf("getkv error, the err is :%s\n", err.Error())
		return sdk.Error("get kv failed")
	}
	unmarshal, err := Unmarshal(value)
	if err != nil {
		fmt.Printf("Unmarshal error, the err is :%s\n", err.Error())
		return sdk.Error("unmarshal value failed")
	}
	m, ok := unmarshal.(*marble)
	if !ok {
		fmt.Printf("the value is not marble\n")
		return sdk.Error("the value is not marble struct1111")
	}
	fmt.Printf("the marble name is %s\n", m.Name)
	return sdk.Success([]byte(m.Name))
}

func initMarble(stubInterface sdk.ContractStub, args [][]byte) common.InvocationResponse {
	const numOfArgs = 2
	if len(args) < numOfArgs {
		fmt.Println("The args number is not correct")
		return sdk.Error("The args is not correct.")
	}
	key := args[0]
	keyStr := string(key)

	value, err := stubInterface.GetKV(keyStr)
	if err != nil {
		errInfo := fmt.Sprintf("Get the marble info err:%s", err.Error())
		return sdk.Error(errInfo)
	}
	if value != nil {
		sdk.Error("The key to be add is already exist")
	}
	value = args[1]

	err = stubInterface.PutKV(keyStr, value)
	if err != nil {
		return sdk.Error(err.Error())
	}
	return sdk.Success(nil)
}

func writeMarble(stubInterface sdk.ContractStub, args [][]byte) common.InvocationResponse {
	key := args[0]
	keyStr := string(key)
	value := args[1]

	err := stubInterface.PutKV(keyStr, value)
	if err != nil {
		return sdk.Error(err.Error())
	}
	return sdk.Success(nil)
}

func getMarble(stubInterface sdk.ContractStub, args [][]byte) common.InvocationResponse {
	if len(args) < 1 {
		fmt.Println("The args number is not correct")
		return sdk.Error("The args is not correct.")
	}

	key := args[0]
	value, err := stubInterface.GetKV(string(key))
	if err != nil {
		errInfo := fmt.Sprintf("Get the key: %s failed", key)
		return sdk.Error(errInfo)
	}

	return sdk.Success(value)
}

func deleteMarble(stubInterface sdk.ContractStub, args [][]byte) common.InvocationResponse {
	if len(args) < 1 {
		return sdk.Error("The args for delete marble is not correct")
	}
	log.Debugf("Enter the delete marble\n")

	key := args[0]
	value, err := stubInterface.GetKV(string(key))
	if err != nil {
		errInfo := fmt.Sprintf("Get the marble info err:%s", err.Error())
		return sdk.Error(errInfo)
	}
	if value == nil {
		sdk.Error("The key to be delete is not exist")
	}
	log.Debugf("Before the delete marble\n")
	err = stubInterface.DelKV(string(key))
	if err != nil {
		return sdk.Error(err.Error())
	}

	return sdk.Success(nil)
}

func initRange(stubInterface sdk.ContractStub, args [][]byte) common.InvocationResponse {
	log.Debugf("Enter initRange\n")

	if len(args) < 1 {
		log.Errorf("The args for initRange is not correct\n")
		return sdk.Error("The args for initRange is not correct")
	}
	numberStr := string(args[0])
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		log.Error("The arg for initRange is not integer\n")
		return sdk.Error("The arg for initRange is not integer")
	}

	log.Debugf("The number is %d\n", number)

	for i := 0; i < number; i++ {
		var key string
		switch {
		case 0 <= i && i <= 9:
			key = fmt.Sprintf("marble00%d", i)
		case 10 <= i && i <= 99:
			key = fmt.Sprintf("marble0%d", i)
		case 100 <= i && i <= 999:
			key = fmt.Sprintf("marble%d", i)
		default:
			return sdk.Error("invalid number")
		}

		value := "white"
		err := stubInterface.PutKV(key, []byte(value))
		if err != nil {
			return sdk.Error(err.Error())
		}
	}
	return sdk.Success(nil)
}

func getRange(stubInterface sdk.ContractStub, args [][]byte) common.InvocationResponse {
	log.Debugf("Enter getRange\n")
	var beginKey string
	var endKey string
	const numOfArgs = 2
	if len(args) < numOfArgs {
		beginKey = ""
		endKey = ""
	} else {
		beginKey = string(args[0])
		endKey = string(args[1])
	}

	iterator, err := stubInterface.GetIterator(beginKey, endKey)
	if err != nil {
		return sdk.Error(err.Error())
	}
	defer iterator.Close()

	rangeMap := make(map[string]string)
	var count = 0
	for {
		b := iterator.Next()
		if b {
			key := iterator.Key()
			value := string(iterator.Value())
			rangeMap[key] = value
			count++
			log.Debugf("The iterator read key is %s, value is %s, count is %d\n", key, value, count)
		} else {
			log.Debugf("The iterator break\n")
			break
		}
	}
	rangeMapBytes, err := json.Marshal(rangeMap)
	if err != nil {
		return sdk.Error(err.Error())
	}
	log.Debugf("rangeMap is %v\n", rangeMap)

	return sdk.Success(rangeMapBytes)
}

func getRangeAndPutKV(stubInterface sdk.ContractStub, args [][]byte) common.InvocationResponse {
	log.Debugf("Enter getRange\n")
	var beginKey string
	var endKey string
	const numOfArgs = 2
	if len(args) < numOfArgs {
		beginKey = ""
		endKey = ""
	} else {
		beginKey = string(args[0])
		endKey = string(args[1])
	}

	iterator, err := stubInterface.GetIterator(beginKey, endKey)
	if err != nil {
		return sdk.Error(err.Error())
	}
	defer iterator.Close()

	rangeMap := make(map[string]string)
	var count = 0
	for {
		b := iterator.Next()
		if b {
			key := iterator.Key()
			value := string(iterator.Value())
			rangeMap[key] = value
			count++
			log.Debugf("The iterator read key is %s, value is %s, count is %d\n", key, value, count)
		} else {
			log.Debugf("The iterator break\n")
			break
		}
	}
	rangeBytes, err := json.Marshal(rangeMap)
	if err != nil {
		return sdk.Error(err.Error())
	}
	log.Debugf("rangeMap is %v\n", rangeMap)

	err = stubInterface.PutKV("jsonMarble", rangeBytes)
	if err != nil {
		return sdk.Error("putstate failed")
	}

	return sdk.Success(rangeBytes)
}

func getPartRangeAndPutKV(stubInterface sdk.ContractStub, args [][]byte) common.InvocationResponse {
	log.Debugf("Enter getPartRange\n")
	var beginKey string
	var endKey string
	const numOfArgs = 2
	if len(args) < numOfArgs {
		beginKey = ""
		endKey = ""
	} else {
		beginKey = string(args[0])
		endKey = string(args[1])
	}

	iterator, err := stubInterface.GetIterator(beginKey, endKey)
	if err != nil {
		return sdk.Error(err.Error())
	}
	defer iterator.Close()

	rangeMap := make(map[string]string)
	var count = 2

	for iterator.Next() {
		key := iterator.Key()
		value := iterator.Value()

		rangeMap[key] = string(value)
		count--
		if count <= 0 {
			break
		}
	}
	log.Debugf("getPartRange is %v\n", rangeMap)
	rangeMapBytes, err := json.Marshal(rangeMap)
	if err != nil {
		return sdk.Error(err.Error())
	}
	err = stubInterface.PutKV("jsonMarble", rangeMapBytes)
	if err != nil {
		return sdk.Error("putstate failed")
	}
	return sdk.Success(nil)
}

func getJSONMarble(stub sdk.ContractStub) common.InvocationResponse {
	value, err := stub.GetKV("jsonMarble")
	if err != nil {
		log.Errorf("Get json marble failed\n")
		return sdk.Error("Get json marble failed")
	}
	log.Infof("The rangeMap is %v\n", value)
	return sdk.Success(value)
}

type marble struct {
	ObjectType string
	Name       string
	Color      string
	Size       int
	Owner      string
}

func (m marble) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal Unmarshal json data.
func Unmarshal(data []byte) (interface{}, error) {
	var value marble
	err := json.Unmarshal(data, &value)
	if err != nil {
		return nil, err
	}
	return &value, nil
}

var storeKey string
var storeValue []byte
var readKey string
var readValue []byte

func getMarbleCom(stubInterface sdk.ContractStub, args [][]byte) common.InvocationResponse {
	if len(args) < 1 {
		fmt.Println("The args number is not correct")
		return sdk.Error("The args is not correct.")
	}
	//fmt.Println("Enter getMarble")
	key := args[0]
	fmt.Printf("getMarbleCom the key is %s\n", string(key))
	value, err := stubInterface.GetKV(string(key))
	if err != nil {
		errInfo := fmt.Sprintf("Get the key: %s failed", key)
		return sdk.Error(errInfo)
	}
	readKey = string(key)
	readValue = value
	compare := bytes.Compare([]byte(storeKey), []byte(readKey))
	if compare == 0 {
		fmt.Printf("key .Compare equal\n")
	}
	i := bytes.Compare(storeValue, readValue)
	if i == 0 {
		fmt.Printf("value .Compare equal\n")
	}

	mb := &marble{}
	err = json.Unmarshal(value, mb)
	if err != nil {
		fmt.Printf("getMarbleCom marble unmarshal failed\n")
		return sdk.Error("getMarbleCom marble unmarshal failed")
	}

	sprintf := fmt.Sprintf("marble name is %s, marble size %d, marble owner %s, "+
		"                          marble color %s, marble object :%s",
		mb.Name, mb.Size, mb.Owner, mb.Color, mb.ObjectType)

	return sdk.Success([]byte(sprintf))
}

func initMarbleAndIndex(stubInterface sdk.ContractStub, args [][]byte) common.InvocationResponse {
	const numOfArgs = 4
	if len(args) < numOfArgs {
		fmt.Printf("The args number for init marble is not correct\n")
		return sdk.Error("The args number for init marble is not correct")
	}
	marbleName := string(args[0])
	marbleOwner := string(args[1])
	marbleColor := string(args[2])
	marbleSizeStr := string(args[3])
	marbleSize, err := strconv.Atoi(marbleSizeStr)
	if err != nil {
		fmt.Printf("The marble size is not int type\n")
		return sdk.Error("The marble size is not int type")
	}

	marbleInfo := &marble{Name: marbleName, Color: marbleColor, Size: marbleSize, Owner: marbleOwner, ObjectType: "marble"}
	marbleBytes, err := json.Marshal(marbleInfo)
	if err != nil {
		fmt.Printf("MarbleInfo marshal error\n")
		return sdk.Error("MarbleInfo marshal error")
	}
	fmt.Printf("The key is marbleInfo.Name:%s\n", marbleInfo.Name)
	err = stubInterface.PutKV(marbleInfo.Name, marbleBytes)
	if err != nil {
		return sdk.Error("putKV error")
	}

	indexName := "colorindex"
	err = stubInterface.SaveComIndex(indexName, []string{marbleInfo.Color}, marbleInfo.Name)
	if err != nil {
		fmt.Printf("SaveComIndex failed\n")
	}

	return sdk.Success(nil)
}

func transferMarblesByIndex(stub sdk.ContractStub, args [][]byte) common.InvocationResponse {
	const numOfArgs = 2
	if len(args) < numOfArgs {
		return sdk.Error("Incorrect number of arguments. Expecting 2")
	}

	color := string(args[0])
	newOwner := strings.ToLower(string(args[1]))
	fmt.Println("- start transferMarblesBasedOnColor ", color, newOwner)

	// Query the color~name index by color
	// This will execute a key range query on all keys starting with 'color'
	kvIterator, err := stub.GetKVByComIndex("colorindex", []string{color})
	if err != nil {
		return sdk.Error(err.Error())
	}

	for {
		hasNext := kvIterator.Next()
		if hasNext {
			key := kvIterator.Key()
			value := kvIterator.Value()
			fmt.Printf("The index has next, final key is %s, value is %s\n", key, string(value))
			marbleToTransfer := marble{}
			err = json.Unmarshal(value, &marbleToTransfer) //unmarshal it aka JSON.parse()
			if err != nil {
				fmt.Printf("Unmarshal marble failed\n")
				return sdk.Error(err.Error())
			}
			marbleToTransfer.Owner = newOwner //change the owner

			marbleJSONasBytes, _ := json.Marshal(marbleToTransfer)
			err := stub.PutKV(key, marbleJSONasBytes)
			if err != nil {
				return sdk.Error(err.Error())
			}
		} else {
			break
		}
	}

	return sdk.Success(nil)
}

func main() {
	smstub.Start(&example01{})
}
