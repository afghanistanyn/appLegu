package cmd

import (
	"fmt"
	"github.com/afghanistanyn/appLegu/utils"
	"os"
)

func Legu(pkgName string, pkgUrl string, pkgMd5 string) {

	conf, err := utils.ReadConf()
	if err != nil {
		fmt.Println("an err occurd on read config file: ", err)
		os.Exit(1)
	}

	client, err := utils.NewMsClient(conf)
	if err != nil {
		fmt.Println("an err occurd on create ms client: ", err)
		os.Exit(1)
	}

	//todo
	//calc md5
	//upload to oss

	//shield
	fmt.Println("start shiled pkg: ", pkgName)
	apkDlUrl, err := utils.ShieldPkg(client, pkgName, pkgUrl, pkgMd5)
	if err != nil {
		fmt.Println("an err occurd on shield pkg: ", err)
		os.Exit(1)
	}

	//download
	fmt.Println("start download shield apk")
	DownloadDestApkFile, err := utils.ApkDownLoad(apkDlUrl, conf.Shield.OutDirectory)
	if err != nil {
		if DownloadDestApkFile != "" {
			os.Remove(DownloadDestApkFile)
		}
		fmt.Println("an err occurd on download the shield apk file: ", err)
		os.Exit(1)
	}

	//align
	fmt.Println("align apk file")
	AlignDestPkg, err := utils.AlignApk(conf, DownloadDestApkFile)
	if err != nil {
		fmt.Println("an err occurd on align the apk file: ", err)
		os.Exit(1)
	}

	//resign
	fmt.Println("resign shield apk file")
	SignedPkg, err := utils.SignApk(conf, "com.zw.cxtpro", AlignDestPkg, true)
	if err != nil {
		fmt.Println("an err occurd on sign the apk file: ", err)
		os.Exit(1)
	}

	//handler result
	fmt.Println("Completion: ", SignedPkg)
}
