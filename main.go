package main

import (
	"fmt"
	"github.com/Insua/database-backup/app/compress"
	"github.com/Insua/database-backup/app/database"
	"github.com/Insua/database-backup/app/file"
	"github.com/Insua/database-backup/app/upyun"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
	"os"
)

func main() {
	sp := g.Cfg().GetString("backup.storage")
	if !gfile.Exists(sp) {
		if err := gfile.Mkdir(sp); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	file.Clean()
	upyun.Clean()

	dir, dbs := database.Init()

	database.Backup(dir, dbs)

	compress.ZipPath(dir)

	upyun.Upload(dir)

	file.DeleteSingle(dir)

}
