package compress

import (
	"fmt"
	"github.com/gogf/gf/encoding/gcompress"
	"os"
)

func ZipPath(dir string)  {
	fmt.Println("compress backup dir")
	if err := gcompress.ZipPath(dir, dir + ".zip"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
