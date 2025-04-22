package sdk

import (
	"context"
	"fmt"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount"
	"github.com/cellargalaxy/go_common/util"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

func GetAppId() string {
	value := util.GetEnv("wx_app_id")
	return value
}
func GetSecret() string {
	value := util.GetEnv("wx_secret")
	return value
}
func GetUsers() []string {
	value := util.GetEnv("wx_user")
	list := strings.Split(value, ",")
	return list
}

var Sdk *officialAccount.OfficialAccount

func InitSdk(ctx context.Context) error {
	var log officialAccount.Log
	log.Stdout = false
	//log.Level = "debug"
	log.Level = "error"
	log.File = "/tmp/wechat_info.log"
	log.Error = "/tmp/wechat_info.log"
	var cfg officialAccount.UserConfig
	cfg.Log = log
	cfg.AppID = GetAppId()
	cfg.Secret = GetSecret()
	logrus.WithContext(ctx).WithFields(logrus.Fields{"cfg": util.ToJsonString(cfg)}).Info("初始化SDK")
	if cfg.AppID == "" || cfg.Secret == "" {
		logrus.WithContext(ctx).WithFields(logrus.Fields{}).Error("初始化SDK，参数为空")
		return fmt.Errorf("初始化SDK，参数为空")
	}
	value, err := officialAccount.NewOfficialAccount(&cfg)
	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"err": err}).Error("初始化SDK，异常")
		return fmt.Errorf("初始化SDK，异常: %v", err)
	}
	Sdk = value
	return nil
}

func GetSdk(ctx context.Context) (*officialAccount.OfficialAccount, error) {
	var err error
	if Sdk == nil {
		err = InitSdk(ctx)
	}
	return Sdk, err
}

func ReTry[Req any, Resp any](ctx context.Context, funz func(ctx context.Context, req Req) (Resp, error), req Req) (resp Resp, err error) {
	for i := 0; i < 5; i++ {
		resp, err = funz(ctx, req)
		if err == nil {
			return resp, nil
		}
		if strings.Contains(err.Error(), "初始化SDK") {
			return resp, err
		}
		util.SleepWare(ctx, time.Second*5)
	}
	return resp, err
}
