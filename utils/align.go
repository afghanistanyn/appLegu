package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

func AlignApk(conf Config, source string) (dstPkg string, err error) {

	apkalign := conf.Shield.ApkAlign
	dstPkg = strings.ReplaceAll(source, ".apk", "_aligned.apk")
	alignCmd := fmt.Sprintf("%s -f 4 align %s %s", apkalign, source, dstPkg)

	fmt.Println(alignCmd)

	cmd := exec.Command(
		apkalign,
		"-f",
		"4",
		source,
		dstPkg,
	)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return source, err
	}
	fmt.Println(string(out))

	return dstPkg, nil
}
