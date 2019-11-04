

#### Shield Apis
```
CreateResourceInstances
CreateShieldInstance
CreateShieldPlanInstance
DeleteShieldInstances
DescribeShieldInstances
DescribeShieldPlanInstance
DescribeShieldResult

```


---
#### Api Address 
    //https://github.com/TencentCloud/tencentcloud-sdk-go/blob/d6fa243ca52014fcde0cbf23efaa9459115581a5/tencentcloud/ms/v20180408/models.go#L243
    ms.ap-guangzhou.tencentcloudapi.com 
---

#### Request Limit 
    GET request should less than 32KB
    POST request signature param should less than 1MB when use HmacSHA1、HmacSHA256 ,less than 10MB when use TC3-HMAC-SHA256.
 ----
 
#### Sign And Align Limit 

    If you are using apksigner, you can only execute zialign before signing the APK file. If you make further changes to the APK after signing it with apkssigner, the signature will be invalidated.
    If you are using jarsigner, you can only execute zialign after signing the APK file.

    
#### Configure

```$yaml
auth:
  txMsSecretId: SecretId
  txMsSecretKey: SecretKey

shield:
  apkSigner: "/usr/local/appLegu/lib/apksigner.jar"       //the path of apksigner
  apkAlign: "/usr/local/appLegu/lib/zipalign"             //the path of zipalign
  apkSigChecker: "/usr/local/appLegu/lib/CheckAndroidV2Signature.jar"       
  outDirectory: "/usr/local/appLegu/pkgs"                //the directory of output apks
  signParams:
    com_zw_cxtpro:                                      //the shield apk bundle name , concat with '_'
      keyAlias: "App"                                  // the sign config of your apk
      keyPassword: "cxtzwcom"
      storeFile: "/usr/local/appLegu/conf/ZWKeystore.jks"
      storePassword: "cxtzwcom"
```


---- 
#### ref:
- https://github.com/TencentCloud/tencentcloud-sdk-go
- https://cloud.tencent.com/document/product/283/17742
- https://developer.android.google.cn/studio/command-line/zipalign.html
----