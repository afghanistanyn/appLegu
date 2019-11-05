
#### install

##### install from source
```$bash

sh package.sh
# now applegu.tar.gz in the top of the source code

# install
# tar vxzf applegu.tar.gz -C /usr/local/

# run 
# /usr/local/applegu/bin/appLegu

```


#### Configure

```$yaml
auth:
  txMsSecretId: SecretId
  txMsSecretKey: SecretKey

shield:
  apkSigner: "/usr/local/applegu/lib/apksigner.jar"       //the path of apksigner
  apkAlign: "/usr/local/applegu/lib/zipalign"             //the path of zipalign
  apkSigChecker: "/usr/local/applegu/lib/CheckAndroidV2Signature.jar"       
  outDirectory: "/usr/local/applegu/pkgs"                //the directory of output apks
  signParams:
    com_zw_cxtpro:                                      //the shield apk bundle name , concat with '_'
      keyAlias: "App"                                  // the sign config of your apk
      keyPassword: "cxtzwcom"
      storeFile: "/usr/local/applegu/conf/ZWKeystore.jks"
      storePassword: "cxtzwcom"
```


---- 
#### ref:
- https://github.com/TencentCloud/tencentcloud-sdk-go
- https://cloud.tencent.com/document/product/283/17742
- https://developer.android.google.cn/studio/command-line/zipalign.html
----