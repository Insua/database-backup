package file

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
	"os"
	"time"
)

func DeleteSingle(dir string)  {
	if gfile.Exists(dir) && gfile.IsDir(dir) {
		gfile.Remove(dir)
	}
	if gfile.Exists(dir+".zip") && gfile.IsFile(dir+".zip") {
		gfile.Remove(dir+".zip")
	}
}

func Clean()  {
	path := g.Cfg().GetString("backup.storage")
	dirs, err := gfile.ScanDir(gfile.RealPath(path), "*")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	for _, v := range dirs {
			info, err := gfile.Info(v)
			if err != nil {
				return
			}
			if time.Now().Sub(info.ModTime()).Hours() > 1 {
				gfile.Remove(v)
			}
	}
}
