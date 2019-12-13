package utils

import (
	"github.com/spf13/viper"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	ms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ms/v20180408"
)

type Config struct {
	Auth   Auth   `yaml:"auth"`
	Shield Shield `yaml:"shiled"`
}

type Auth struct {
	TxMsSecretId  string `yaml:"txMsSecretId"`
	TxMsSecretKey string `yaml:"txMsSecretKey"`
}

type Shield struct {
	ApkSigner     string                   `yaml:"jarSigner"`
	ApkAlign      string                   `yaml:"apkAlign"`
	ApkSigChecker string                   `yaml:"apkSigChecker"`
	OutDirectory  string                   `yaml:"outDirectory"`
	ShieldTimeout uint16                   `yaml:"shieldTimeout"`
	Signparams    map[string]*AppSignParam `yaml:"signParams"`
}

type AppSignParam struct {
	KeyAlias      string `yaml:"keyAlias"`
	KeyPassword   string `yaml:"keyPassword"`
	StoreFile     string `yaml:"storeFile"`
	StorePassword string `yaml:"storePassword"`
}

func ReadConf() (appConf Config, err error) {

	myconf := Config{}
	conf := viper.New()
	conf.AddConfigPath("./")
	conf.AddConfigPath("conf/")
	conf.AddConfigPath("../conf/")
	conf.AddConfigPath("/etc/applegu/")
	conf.AddConfigPath("/usr/local/applegu/conf/")
	conf.AddConfigPath("/usr/local/etc/")
	conf.AddConfigPath("/usr/local/applegu/etc/")
	conf.AddConfigPath("/etc/applegu/")
	conf.SetConfigType("yaml")
	if err := conf.ReadInConfig(); err != nil {
		return myconf, err
	}

	err = conf.Unmarshal(&myconf)
	if err != nil {
		return myconf, err
	}

	return myconf, nil
}

func NewMsClient(conf Config) (client *ms.Client, err error) {
	credential := common.NewCredential(conf.Auth.TxMsSecretId, conf.Auth.TxMsSecretKey)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqTimeout = 30
	cpf.Language = "en-US"

	return ms.NewClient(credential, regions.Guangzhou, cpf)
}
