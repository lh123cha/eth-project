/*
 * Copyright (c) Huawei Technologies Co., Ltd. 2020-2020. All rights reserved.
 */

package contract

import (
	"fmt"
	"strings"
	"time"

	"git.huawei.com/poissonsearch/wienerchain/wienerchain-go-sdk/node"

	"git.huawei.com/poissonsearch/wienerchain/proto/common"
	"git.huawei.com/poissonsearch/wienerchain/proto/nodeservice"
	"demo/utils"
	t "git.huawei.com/poissonsearch/wienerchain/wienerchain-go-sdk/utils"

	"git.huawei.com/poissonsearch/wienerchain/wienerchain-go-sdk/client"
	"github.com/golang/protobuf/proto"
	"github.com/spf13/cobra"
)

const (
	// WaitTime is used to set the wait time for update chain.
	WaitTime = 60
)

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "send transaction commands",
	Long:  "send transaction commands",
	Run: func(cmd *cobra.Command, args []string) {
		send()
	},
}

func getSendCmd() *cobra.Command {
	bindFlags(sendCmd)
	sendCmd.Flags().StringVarP(&nodes, "nodes", "n", "", "please specify nodes.")
	if err := sendCmd.MarkFlagRequired("nodes"); err != nil {
		fmt.Printf("mark flag required error: %v\n", err)
	}
	if err := sendCmd.MarkFlagRequired("function"); err != nil {
		fmt.Printf("mark flag required error: %v\n", err)
	}
	if err := sendCmd.MarkFlagRequired("args"); err != nil {
		fmt.Printf("mark flag required error: %v\n", err)
	}
	return sendCmd
}

func send() {
	args = strings.TrimSpace(args)
	args := strings.Split(args, ";")
	sendTransaction(args)
}

func sendTransaction(args []string) {
	var err error
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

	rawMsg, err := gatewayClient.ContractRawMessage.BuildInvokeMessage(chainID, txID, contract, function, args)
	if err != nil {
		fmt.Printf("contract raw messessage build invoke message error： %v", err)
		return
	}
	nodeNames := strings.Split(nodes, ",")
	nodeMap := gatewayClient.Nodes

	var invokeResponses []*common.RawMessage
	for _, nodeName := range nodeNames {
		node, ok := nodeMap[nodeName]
		if !ok {
			fmt.Printf("node not exist： %v\n", nodeName)
			return
		}
		var invokeResponse *common.RawMessage
		invokeResponse, err = node.ContractAction.Invoke(rawMsg)
		if err != nil {
			fmt.Printf("invoke error： %v", err)
			return
		}
		invokeResponses = append(invokeResponses, invokeResponse)
	}

	transactionRawMsg, err := gatewayClient.ContractRawMessage.BuildTransactionMessage(invokeResponses)
	if err != nil {
		fmt.Printf("build transaction message error: %v\n", err)
		return
	}

	processResult(transactionRawMsg, nodeMap, txID, nodeNames[0])
}

func processResult(transactionRawMsg *common.RawMessage, nodeMap map[string]*node.WNode, txID string,
	listenNodeName string) {
	node, ok := nodeMap[listenNodeName]
	if !ok {
		fmt.Printf("node not exist： %v\n", listenNodeName)
		return
	}
	event, err := node.EventAction.Listen(chainID)
	if err != nil {
		fmt.Printf("event action listen error: %v", err)
		return
	}
	in := make(chan string)
	go listen(event, txID, in)
	if consenters == "" {
		consenters = nodes
	}
	consenterNodeNames := strings.Split(consenters, ",")
	var transactionResponses []*common.Response
	for _, nodeName := range consenterNodeNames {
		node, ok := nodeMap[nodeName]
		if !ok {
			fmt.Printf("node not exist： %v\n", nodeName)
			return
		}
		transactionResponse, err := node.ContractAction.Transaction(transactionRawMsg)
		if err != nil {
			fmt.Printf("invoke error： %v\n", err)
			return
		}
		txResponse := &common.Response{}
		if err := proto.Unmarshal(transactionResponse.Payload, txResponse); err != nil {
			fmt.Printf("unmarshal transaction response error： %v\n", err)
			return
		}
		transactionResponses = append(transactionResponses, txResponse)
	}
	select {
	case recv := <-in:
		for i, response := range transactionResponses {
			if response.Status == common.Status_SUCCESS {
				utils.PrintResponse(response, consenterNodeNames[i], recv)
			} else {
				utils.PrintResponse(response, consenterNodeNames[i], "")
			}
		}
		return
	case <-time.After(WaitTime * time.Second):
		fmt.Printf("Invoke time out\n")
		return
	}
}

func listen(event nodeservice.EventService_RegisterBlockEventClient, txID string, in chan string) {
	for {
		responseMsg, err := event.Recv()
		if err != nil {
			fmt.Printf("event receive response message error: %v", err)
			return
		}
		res := &common.BlockResult{}
		err = proto.Unmarshal(responseMsg.Payload, res)
		if err != nil {
			fmt.Printf("unmarshal block result error: %v", err)
			return
		}

		for _, txRes := range res.TxResults {
			if txRes.TxId == txID {
				in <- txRes.Status.String()
				return
			}
		}
	}
}
