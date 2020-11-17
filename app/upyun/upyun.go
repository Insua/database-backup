package upyun

import (
	"github.com/gogf/gf/frame/g"
	"github.com/upyun/go-sdk/upyun"
)

func Upyun() *upyun.UpYun {
	return  upyun.NewUpYun(&upyun.UpYunConfig{
		Bucket:   g.Cfg().GetString("backup.upyun.bucket"),
		Operator: g.Cfg().GetString("backup.upyun.operator"),
		Password: g.Cfg().GetString("backup.upyun.password"),
	})
}
