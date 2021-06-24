/*
 * Copyright (c) Huawei Technologies Co., Ltd. 2020-2020. All rights reserved.
 */

// Package contract is the implementation of contract operation for go client.
package contract

import (
	"fmt"
	"github.com/spf13/cobra"
)

var nodes string
var nodeName string
var consenters string
var chainID string
var contract string
var function string
var args string
var configPath string

var contractCmd = &cobra.Command{
	Use:   "contract",
	Short: "contract commands",
	Long:  "contract commands",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

// Cmd is the entrance for contract command.
func Cmd() *cobra.Command {
	contractCmd.AddCommand(getSendCmd())
	//contractCmd.AddCommand(getInitCmd())
	contractCmd.AddCommand(getQueryCmd())
	return contractCmd
}

func bindFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&consenters, "consenters", "s", "", "please specify consenters to send, separate with ','.")
	cmd.Flags().StringVarP(&chainID, "chainId", "c", "", "please specify chain id.")
	if err := cmd.MarkFlagRequired("chainId"); err != nil {
		fmt.Printf("mark flag required error: %v\n", err)
	}
	cmd.Flags().StringVarP(&contract, "contract", "t", "", "please specify contract name.")
	if err := cmd.MarkFlagRequired("contract"); err != nil {
		fmt.Printf("mark flag required error: %v\n", err)
	}
	cmd.Flags().StringVarP(&function, "function", "f", "", "please specify function name.")
	cmd.Flags().StringVarP(&args, "args", "a", "", "please specify args of function, separate with ' '.")
	cmd.Flags().StringVarP(&configPath, "config", "g", "", "please config file path.")
	if err := cmd.MarkFlagRequired("config"); err != nil {
		fmt.Printf("mark flag required error: %v\n", err)
	}
}
