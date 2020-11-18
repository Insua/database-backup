package compress

import (
	"fmt"
	"github.com/gogf/gf/encoding/gcompress"
	"github.com/gogf/gf/os/gtime"
	"os"
	"runtime"
)

func ZipPath(dir string)  {
	defer runtime.GC()
	fmt.Println(gtime.Now().Format("Y-m-d H:i:s") + "- compress backup dir")
	if err := gcompress.ZipPath(dir, dir + ".zip"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(gtime.Now().Format("Y-m-d H:i:s") + "- compress backup complete")
}
