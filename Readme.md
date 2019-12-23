
#### Install

##### install from source
```bash

#run package.sh to build and package applegu
sh package.sh

# install
tar vxzf applegu.tar.gz -C /usr/local/

# run 
/usr/local/applegu/bin/appLegu

```

##### install from pre-build package
```bash
wget -O applegu-latest.tar.gz $(curl -s https://api.github.com/repos/afghanistanyn/applegu/releases/latest | jq -r ".assets[] | .browser_download_url")
tar vxzf applegu-latest.tar.gz -C /usr/local/
```


#### Configure

    vim /usr/local/applegu/conf/config.yaml

```yaml
auth:
  txMsSecretId: SecretId
  txMsSecretKey: SecretKey

shield:
  apkSigner: "/usr/local/applegu/lib/apksigner.jar"       #the path of apksigner
  apkAlign: "/usr/local/applegu/lib/zipalign"             #the path of zipalign
  apkSigChecker: "/usr/local/applegu/lib/CheckAndroidV2Signature.jar"       
  outDirectory: "/usr/local/applegu/pkgs"                #the directory of output apks
  shieldTimeout: 1800                                    #time for wait legu shield 
  checkInterval: 30                                      #interval of shield check , unit sec
  signParams:
    com_zw_cxtpro:                                      #the shield apk bundle name , concat with '_'
      keyAlias: "App"                                   #the sign config of your apk
      keyPassword: "cxtzwcom"
      storeFile: "/usr/local/applegu/conf/ZWKeystore.jks"
      storePassword: "cxtzwcom"
```

#### Run 

    # legu apk (include resign)
    /usr/local/applegu/bin/appLegu legu --pkgmd5 "xxx" --pkgname "xxx" --pkgurl "xxx"

    #resign apk (resign only)
    /usr/local/applegu/bin/appLegu sign --srcpkg "xxx"  --removealign "false"
    
    #check the shield status
    /usr/local/applegu/bin/appLegu check --itemid "xxx" --interval 5 --count 10 --untilsuccess "false"
    
---- 

#### Docker support
```
docker pull afghanistanyn/applegu:latest
```

test
```
docker run -it --rm afghanistanyn/applegu
```

mount conf and keystore
```
docker run -it --rm  -v /app/applegu/config.yaml:/usr/local/applegu/conf/config.yaml -v /app/applegu/xxx.jks:/usr/local/applegu/conf/xxx.jks afghanistanyn/applegu
```

mount output dir
```
docker run -it --rm  -v /app/applegu/config.yaml:/usr/local/applegu/conf/config.yaml -v /app/applegu/xxx.jks:/usr/local/applegu/conf/xxx.jks -v /app/applegu/output:/usr/local/applegu/pkgs afghanistanyn/applegu
```

----

#### Ref:
- https://github.com/TencentCloud/tencentcloud-sdk-go
- https://cloud.tencent.com/document/product/283/17742
- https://developer.android.google.cn/studio/command-line/zipalign.html
----