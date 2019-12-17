/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"github.com/spf13/cobra"
)

var itemId string
var count, interval int
var untilsuccess bool

// signCmd represents the sign command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "check the shield process",

	Args: func(cmd *cobra.Command, args []string) error {

		if len(itemId) == 0 {
			return errors.New("itemid is required")
		}

		if len(itemId) < 32 {
			return errors.New("itemid in error format")
		}

		if count < 0 {
			return errors.New("check count > 0")
		}

		if interval < 0 {
			return errors.New("check interval > 0")
		}

		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {
		Check(itemId, count, interval, untilsuccess)
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
	checkCmd.Flags().StringVarP(&itemId, "itemid", "i", "", "the id of shield process")
	checkCmd.Flags().IntVarP(&count, "count", "c", 1, "the count of check shield process, if success the process will exit")
	checkCmd.Flags().IntVarP(&interval, "interval", "t", 30, "the interval of check shield process, unit sec")
	checkCmd.Flags().BoolVarP(&untilsuccess, "untilsuccess", "f", false, "check the shield process until success, the count will be ignore if set")
}
