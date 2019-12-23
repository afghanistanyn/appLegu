package utils

import (
	"fmt"
	"github.com/shogo82148/androidbinary/apk"
	"os/exec"
	"strings"
)

func SignApk(conf Config, source string, checkv2 bool) (signDestPkg string, err error) {

	//parse pkgName
	fmt.Println("SignApk, parse apk PackageName")
	pkg, err := apk.OpenFile(source)
	if err != nil {
		return "", err
	}
	appPkgName := pkg.PackageName()
	defer func() {
		if ferr := pkg.Close(); ferr != nil {
			err = ferr
		}
	}()

	//the apkName in config file should be concat by '_'
	appPkgNameKey := strings.Replace(appPkgName, ".", "_", -1)
	appSignParam := conf.Shield.Signparams[appPkgNameKey]
	if appSignParam == nil {
		return "", fmt.Errorf("Sign params for %s not found in configure file", appPkgName)
	}

	apksigner := conf.Shield.ApkSigner

	StoreFile := appSignParam.StoreFile
	StorePassword := appSignParam.StorePassword
	KeyPassword := appSignParam.KeyPassword
	KeyAlias := appSignParam.KeyAlias

	dstPkg := strings.ReplaceAll(source, ".apk", "_signed.apk")

	//the java shold find in PATH
	signCmd := fmt.Sprintf("java -jar %s sign --in %s --out %s --ks %s --ks-key-alias %s --ks-pass pass:%s -key-pass pass:%s --v1-signing-enabled --v2-signing-enabled --verbose",
		apksigner,
		source,
		dstPkg,
		StoreFile,
		KeyAlias,
		StorePassword,
		KeyPassword,
	)
	fmt.Println(signCmd)

	cmd := exec.Command(
		"java",
		"-jar", apksigner,
		"sign",
		"--in", source,
		"--out", dstPkg,
		"--ks", StoreFile,
		"--ks-key-alias", KeyAlias,
		"--ks-pass", "pass:"+StorePassword,
		"--key-pass", "pass:"+KeyPassword,
		"--v1-signing-enabled",
		"--v2-signing-enabled",
		"--verbose",
	)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	fmt.Println(string(out))

	if checkv2 {
		fmt.Println("verify the sign")
		apksigchecker := conf.Shield.ApkSigChecker
		cmd := exec.Command(
			"java",
			"-jar", apksigchecker,
			dstPkg,
		)

		out, err := cmd.CombinedOutput()
		if err != nil {
			return "", err
		}
		fmt.Println(string(out))
	}

	return dstPkg, nil
}
