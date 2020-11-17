package upyun

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
	"github.com/upyun/go-sdk/upyun"
	"time"
)

func Clean()  {
	path := g.Cfg().GetString("backup.upyun.folder")
	keepDays := g.Cfg().GetFloat64("backup.upyun.keepDays")
	if keepDays == 0 {
		keepDays = 7
	}
	up := Upyun()
	listChan := make(chan *upyun.FileInfo, 100)
	go func() {
		up.List(&upyun.GetObjectsConfig{
			Path:        path,
			ObjectsChan: listChan,
		})
	}()

	for l := range listChan {
		if time.Now().Sub(l.Time).Hours() / 24 > keepDays {
			up.Delete(&upyun.DeleteObjectConfig{
				Path:  gfile.Join(path, l.Name),
			})
		}
	}
}
