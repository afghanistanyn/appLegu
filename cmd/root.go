package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/url"
	"os"
)

func init() {
	rootCmd.Flags().StringP("pkgName", "n", "", "Bundle name of you apk")
	rootCmd.Flags().StringP("pkgUrl", "u", "", "Download url of you apk, without auth (required)")
	rootCmd.Flags().StringP("pkgMd5", "m", "", "Md5 hash of you apk (required)")

	viper.BindPFlag("pkgName", rootCmd.Flags().Lookup("pkgName"))
	viper.BindPFlag("pkgUrl", rootCmd.Flags().Lookup("pkgUrl"))
	viper.BindPFlag("pkgMd5", rootCmd.Flags().Lookup("pkgMd5"))
}

var rootCmd = &cobra.Command{
	Use:   "appLegu",
	Short: "Used to shield apk by legu",
	//Args: cobra.MinimumNArgs(1),
	//Args: cobra.OnlyValidArgs,

	Args: func(cmd *cobra.Command, args []string) error {
		pkgurlp, _ := cmd.Flags().GetString("pkgUrl")
		pkgmd5p, _ := cmd.Flags().GetString("pkgMd5")
		pkgnamep, _ := cmd.Flags().GetString("pkgName")

		if len(pkgnamep) == 0 {
			return errors.New("pkgName is required , like 'com.tx.webchat'")
		}

		if len(pkgmd5p) != 32 {
			if len(pkgmd5p) == 0 {
				return errors.New("pkgMd5 is required")
			}
			return errors.New("pkgMd5 with incorrect length")
		}
		if len(pkgurlp) == 0 {
			return errors.New("pkgUrl is required")
		}
		_, err := url.Parse(pkgurlp)
		if err != nil {
			return errors.New("pkgUrl with incorrect format")
		}
		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {

		pkgName := viper.GetString("pkgName")
		pkgUrl := viper.GetString("pkgUrl")
		pkgMd5 := viper.GetString("pkgMd5")

		Legu(pkgName, pkgUrl, pkgMd5)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
