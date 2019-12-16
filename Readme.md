
#### Install

##### install from source
```bash

#run package.sh to build and package applegu
sh package.sh

# install & upgrade
tar vxzf applegu.tar.gz -C /usr/local/
chmod u+x /usr/local/applegu/lib/zipalign

# run 
# /usr/local/applegu/bin/appLegu

```

####Docker support
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