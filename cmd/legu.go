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
	"net/url"
)

var pkgName, pkgUrl, pkgMd5 string

var leguCmd = &cobra.Command{
	Use:   "legu",
	Short: "Used to shield apk by legu",

	Args: func(cmd *cobra.Command, args []string) error {

		if len(pkgName) == 0 {
			return errors.New("pkgname is required , like 'com.tx.webchat'")
		}

		if len(pkgMd5) != 32 {
			if len(pkgMd5) == 0 {
				return errors.New("pkgmd5 is required")
			}
			return errors.New("pkgmd5 with incorrect length, 32 characters required")
		}
		if len(pkgUrl) == 0 {
			return errors.New("pkgurl is required")
		}
		_, err := url.Parse(pkgUrl)
		if err != nil {
			return errors.New("pkgurl with incorrect format")
		}
		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {
		Legu(pkgName, pkgUrl, pkgMd5)
	},
}

func init() {
	rootCmd.AddCommand(leguCmd)
	leguCmd.Flags().StringVar(&pkgName, "pkgname", "", "the bundle name of apk file")
	leguCmd.Flags().StringVar(&pkgUrl, "pkgurl", "", "the url of apk file , should be access by tencent ms")
	leguCmd.Flags().StringVar(&pkgMd5, "pkgmd5", "", "the md5sum of apk file , 32 characters required")
}
