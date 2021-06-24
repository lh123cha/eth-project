/*
 * Copyright (c) Huawei Technologies Co., Ltd. 2020-2020. All rights reserved.
 */

package contract

import (
	"fmt"
	"strings"

	"demo/utils"
	t "git.huawei.com/poissonsearch/wienerchain/wienerchain-go-sdk/utils"

	"git.huawei.com/poissonsearch/wienerchain/proto/common"
	"github.com/golang/protobuf/proto"

	"git.huawei.com/poissonsearch/wienerchain/wienerchain-go-sdk/client"

	"github.com/spf13/cobra"
)

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "query commands",
	Long:  "query commands",
	Run: func(cmd *cobra.Command, args []string) {
		query()
	},
}

func getQueryCmd() *cobra.Command {
	bindFlags(queryCmd)
	queryCmd.Flags().StringVarP(&nodeName, "node", "n", "", "please specify node.")
	if err := queryCmd.MarkFlagRequired("node"); err != nil {
		fmt.Printf("mark flag required error: %v\n", err)
	}
	if err := queryCmd.MarkFlagRequired("function"); err != nil {
		fmt.Printf("mark flag required error: %v\n", err)
	}
	if err := queryCmd.MarkFlagRequired("args"); err != nil {
		fmt.Printf("mark flag required error: %v\n", err)
	}
	return queryCmd
}

func query() {
	gatewayClient, err := client.NewGatewayClient(configPath)
	if err != nil {
		fmt.Printf("new gateway client error: %v", err)
		return
	}
	txID, err := t.GenerateTxID()
	if err != nil {
		fmt.Printf("generate tx id error: %v", err)
		return
	}
	args := strings.Split(args, ";")
	rawMsg, err := gatewayClient.ContractRawMessage.BuildInvokeMessage(chainID, txID, contract, function, args)
	if err != nil {
		fmt.Printf("contract raw messessage build vote message error： %v", err)
		return
	}

	nodeMap := gatewayClient.Nodes
	node, ok := nodeMap[nodeName]
	if !ok {
		fmt.Printf("node not exist： %v\n", nodeName)
		return
	}
	invokeResponse, err := node.ContractAction.Invoke(rawMsg)
	if err != nil {
		fmt.Printf("invoke error： %v", err)
		return
	}
	processQueryResult(invokeResponse)
}

func processQueryResult(invokeResponse *common.RawMessage) {
	response := &common.Response{}
	if err := proto.Unmarshal(invokeResponse.Payload, response); err != nil {
		fmt.Printf("unmarshal invoke response error： %v", err)
	}
	if response.Status == common.Status_SUCCESS {
		tx := &common.Transaction{}
		if err := proto.Unmarshal(response.Payload, tx); err != nil {
			fmt.Printf("unmarshal transaction error: %v\n", err)
			return
		}
		txPayLoad := &common.TxPayload{}
		if err := proto.Unmarshal(tx.Payload, txPayLoad); err != nil {
			fmt.Printf("unmarshal tx payload error: %v\n", err)
			return
		}
		txData := &common.CommonTxData{}
		if err := proto.Unmarshal(txPayLoad.Data, txData); err != nil {
			fmt.Printf("unmarshal common tx data error: %v\n", err)
			return
		}
		utils.PrintResponse(response, nodeName, string(txData.Response.Payload))
	} else {
		utils.PrintResponse(response, nodeName, "")
	}
}
