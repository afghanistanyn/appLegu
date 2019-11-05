
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

    If you are using apksigner, you can only execute zipalign before signing the APK file. If you make further changes to the APK after signing it with apkssigner, the signature will be invalidated.
    If you are using jarsigner, you can only execute zipalign after signing the APK file.
    