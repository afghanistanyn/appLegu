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
	"os"
)

var srcPkg string
var removeAlign bool

// signCmd represents the sign command
var signCmd = &cobra.Command{
	Use:   "sign",
	Short: "only sign the apk file",

	Args: func(cmd *cobra.Command, args []string) error {

		if len(srcPkg) == 0 {
			return errors.New("srcpkg is required")
		}

		_, err := os.Stat(srcPkg)
		if err != nil {
			if !os.IsExist(err) {
				return errors.New("srcpkg not exist")
			}
		}

		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {
		Sign(srcPkg, removeAlign)
	},
}

func init() {
	rootCmd.AddCommand(signCmd)
	signCmd.Flags().StringVar(&srcPkg, "srcpkg", "", "the path of sign pkg")
	signCmd.Flags().BoolVarP(&removeAlign, "removealign", "r", true, "remove the aligned apk")
}
