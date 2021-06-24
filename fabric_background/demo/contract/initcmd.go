/*
 * Copyright (c) Huawei Technologies Co., Ltd. 2020-2020. All rights reserved.
 */

package contract

import (
	"fmt"
	"strings"
	"github.com/spf13/cobra"
)

const (
	// FUNCTION is the flag for init contract.
	FUNCTION = "init"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init contract commands",
	Long:  "init contract commands",
	Run: func(cmd *cobra.Command, args []string) {
		initContract()
	},
}

func getInitCmd() *cobra.Command {
	bindFlags(initCmd)
	initCmd.Flags().StringVarP(&nodes, "nodes", "n", "", "please specify nodes.")
	if err := initCmd.MarkFlagRequired("nodes"); err != nil {
		fmt.Printf("mark flag required error: %v\n", err)
	}
	if err := initCmd.MarkFlagRequired("consenters"); err != nil {
		fmt.Printf("mark flag required error: %v\n", err)
	}
	if err := initCmd.MarkFlagRequired("args"); err != nil {
		fmt.Printf("mark flag required error: %v\n", err)
	}
	return initCmd
}

func initContract() {
	function = FUNCTION
	args = strings.TrimSpace(args)
	args := strings.Split(args, ";")
	sendTransaction(args)
}
