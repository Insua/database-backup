package upyun

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/gtime"
	"github.com/upyun/go-sdk/upyun"
	"os"
)


func Upload(dir string)  {
	up := Upyun()
	if err := up.Put(&upyun.PutObjectConfig{
		Path:      gfile.Join(g.Cfg().GetString("backup.upyun.folder"), gtime.Now().Format("Y-m-d-H-i-s") + ".zip"),
		LocalPath: dir + ".zip",
	}); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
