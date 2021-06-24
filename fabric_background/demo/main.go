/*
 * Copyright (c) Huawei Technologies Co., Ltd. 2020-2020. All rights reserved.
 */

// Package main is the entrance of go client.
package main

import (
	"os"
	"github.com/spf13/cobra"
	"demo/contract"
)

func newMainCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "wClient",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},

		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
	rootCmd.AddCommand(contract.Cmd())
	return rootCmd
}

// Execute is the entrance of main command.
func Execute() {
	if err := newMainCmd().Execute(); err != nil {
		os.Exit(1)
	}
}

func main() {
	Execute()
}
