package utils

import (
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	mserrors "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	ms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ms/v20180408"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"time"
)

func GetAppSetInfoByPkgName(client *ms.Client, appPkgName string, appVersion string, itemId string) (appSetInfo *ms.AppSetInfo, err error) {

	request := ms.NewDescribeShieldInstancesRequest()
	request.Filters = []*ms.Filter{
		{
			Name:  common.StringPtr("AppPkgName"),
			Value: common.StringPtr(appPkgName),
		},
	}

	if appVersion != "" {
		versionFilter := &ms.Filter{
			Name:  common.StringPtr("AppVersion"),
			Value: common.StringPtr(appVersion),
		}
		request.Filters = append(request.Filters, versionFilter)
	}

	if itemId != "" {
		itemIdFilter := &ms.Filter{
			Name:  common.StringPtr("ItemId"),
			Value: common.StringPtr(itemId),
		}
		request.Filters = append(request.Filters, itemIdFilter)
	}

	response, err := client.DescribeShieldInstances(request)
	if err != nil {
		appSetInfo = &ms.AppSetInfo{
			AppPkgName: &appPkgName,
			AppVersion: &appVersion,
		}
		return appSetInfo, err
	}

	//for debug
	fmt.Println(response.ToJsonString())

	if len(response.Response.AppSet) == 0 {
		return appSetInfo, nil
	}
	return response.Response.AppSet[0], nil
}

func CheckShield(client *ms.Client, itemId string) {

	req := ms.NewDescribeShieldResultRequest()
	req.ItemId = &itemId
	resultResp, err := client.DescribeShieldResult(req)
	if err != nil {
		fmt.Printf("error occurd when get shield result of itemid %s: %s", itemId, err)
		os.Exit(-1)
	}

	fmt.Println(resultResp.ToJsonString())
}

func ShieldPkg(client *ms.Client, pkgName string, pkgUrl string, pkgMd5 string, waitTime uint16) (apkDlUrl string, err error) {

	req := ms.NewCreateShieldInstanceRequest()
	appInfo := ms.AppInfo{
		AppName: &pkgName,
		AppUrl:  &pkgUrl,
		AppMd5:  &pkgMd5,
	}

	serviceInfo := ms.ServiceInfo{
		ServiceEdition: common.StringPtr("basic"),
		SubmitSource:   common.StringPtr("RDM-rdm"),
		CallbackUrl:    common.StringPtr(""),
	}

	req.AppInfo = &appInfo
	req.ServiceInfo = &serviceInfo

	shieldResp, err := client.CreateShieldInstance(req)
	if err != nil {
		return "", err
	}

	//for debug
	fmt.Println(shieldResp.ToJsonString())

	itemId := shieldResp.Response.ItemId

	// 任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
	//TaskStatus *uint64 `json:"TaskStatus,omitempty" name:"TaskStatus"`

	//get shield info every 30s
	during := 30
	if int(waitTime) < during {
		waitTime = uint16(during)
	}
	retry_count := int(waitTime / uint16(during))
	for i := 0; i <= retry_count; i++ {
		shieldSetInfo, _ := GetAppSetInfoByPkgName(client, "", "", *itemId)

		//fmt.Println(*shieldSetInfo.TaskStatus)
		if *shieldSetInfo.TaskStatus == 2 {
			//the last retry
			if i == retry_count {
				return "", mserrors.NewTencentCloudSDKError(string(*shieldSetInfo.TaskStatus), "shield timeout", *shieldSetInfo.ItemId)
			}
			fmt.Println("shielding ...")
			time.Sleep(time.Duration(during) * time.Second)
		}

		if *shieldSetInfo.TaskStatus == 1 {
			fmt.Println("shiled completion ...")
			return *shieldSetInfo.AppUrl, nil
		}

		if *shieldSetInfo.TaskStatus == 3 || *shieldSetInfo.TaskStatus == 4 {
			fmt.Println("shiled error ...")
			return "", mserrors.NewTencentCloudSDKError(string(*shieldSetInfo.TaskStatus), "shiled error or timeout", *shieldSetInfo.ItemId)
		}
	}
	return "", nil
}

func ApkDownLoad(ApkUrl string, OutDirectory string) (dstFilePath string, err error) {

	u, err := url.Parse(ApkUrl)
	if err != nil {
		return "", err
	}

	paths := strings.Split(u.Path, "/")
	dstFile := paths[len(paths)-1]
	dstFilePath = path.Join(OutDirectory, dstFile)

	if !exists(OutDirectory) {
		err := os.MkdirAll(OutDirectory, os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	dstApkFile, err := os.Create(dstFilePath)
	if err != nil {
		return dstFilePath, err
	}

	defer func() {
		if ferr := dstApkFile.Close(); ferr != nil {
			err = ferr
		}
	}()

	resp, err := http.Get(ApkUrl)
	if err != nil {
		return dstFilePath, err
	}
	defer func() {
		if ferr := resp.Body.Close(); ferr != nil {
			err = ferr
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return dstFilePath, fmt.Errorf("download apk file error http_status: %s", resp.Status)
	}

	_, err = io.Copy(dstApkFile, resp.Body)
	if err != nil {
		return dstFilePath, err
	}
	return dstFilePath, err
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
