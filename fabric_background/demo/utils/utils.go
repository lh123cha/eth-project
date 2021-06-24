/*
 * Copyright (c) Huawei Technologies Co., Ltd. 2020-2020. All rights reserved.
 */

// Package utils is the implementation of some common functions.
package utils

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	"git.huawei.com/poissonsearch/wienerchain/proto/common"
)

// PrintResponse is used to print response.
func PrintResponse(response *common.Response, node string, info string) {
	if response.Status == common.Status_SUCCESS {
		PrintInfo(node, response.Status.String(), info)
	} else {
		PrintInfo(node, response.Status.String(), response.StatusInfo)
	}
}

// PrintInfo is used to print info with json format.
func PrintInfo(node string, status string, info string) {
	infoMap := make(map[string]interface{})
	infoMap["info"] = info
	infoMap["status"] = status
	infoMap["node"] = node
	b, err := json.Marshal(infoMap)
	if err != nil {
		fmt.Printf("marshal json error: %v", err)
	}
	fmt.Println(string(b))
}

// Hash2str is used to convert the hash value to string for human reading.
func Hash2str(h []byte) string {
	return hex.EncodeToString(h)
}
