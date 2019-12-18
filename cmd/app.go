package cmd

import (
	"fmt"
	"github.com/afghanistanyn/appLegu/utils"
	"os"
	"time"
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

	//shield
	fmt.Println("start shiled pkg: ", pkgName)
	waitTime := conf.Shield.ShieldTimeout
	checkInterval := conf.Shield.CheckInterval

	if checkInterval == 0 {
		checkInterval = 30
	}

	apkDlUrl, err := utils.ShieldPkg(client, pkgName, pkgUrl, pkgMd5, waitTime, checkInterval)
	if err != nil {
		fmt.Println("an err occurd on shield pkg: ", err)
		os.Exit(1)
	}

	//download
	fmt.Println("start download shield apk")
	DownloadDestApkFile, err := utils.ApkDownLoad(apkDlUrl, conf.Shield.OutDirectory)
	if err != nil {
		if DownloadDestApkFile != "" {
			_ = os.Remove(DownloadDestApkFile)
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
	fmt.Println("Legu Shield Completion: ", SignedPkg)
}

func Sign(srcPkg string, removeAlign bool) {

	conf, err := utils.ReadConf()
	if err != nil {
		fmt.Println("an err occurd on read config file: ", err)
		os.Exit(1)
	}

	//align
	fmt.Println("align apk file")
	AlignDestPkg, err := utils.AlignApk(conf, srcPkg)
	if err != nil {
		fmt.Println("an err occurd on align the apk file: ", err)
		os.Exit(1)
	}

	//resign
	fmt.Println("resign shield apk file")
	SignedPkg, err := utils.SignApk(conf, "com.zw.cxtpro", AlignDestPkg, true)
	if err != nil {
		fmt.Println("an err occurd on sign the apk file: ", err)
		_ = os.Remove(AlignDestPkg)
		os.Exit(1)
	}

	if removeAlign {
		_ = os.Remove(AlignDestPkg)
	}

	//handler result
	fmt.Println("Sign Completion: ", SignedPkg)
}

//Check(itemId,count,interval,untilsuccess)
func Check(itemId string, count int, interval int, untilsuccess bool) {
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

	if untilsuccess {
		for {
			success := utils.CheckShield(client, itemId)
			if success {
				break
			}
			time.Sleep(time.Duration(interval) * time.Second)
		}
	} else {
		for i := 0; i < count; i++ {
			success := utils.CheckShield(client, itemId)
			if success {
				break
			}
			time.Sleep(time.Duration(interval) * time.Second)
		}
	}
}
