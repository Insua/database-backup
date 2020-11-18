package database

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/gproc"
	"github.com/gogf/gf/os/gtime"
)

func Backup(dir string, dbs []string)  {
	for _, v := range dbs {
		db(dir, v)
	}
}

func db(dir, db string)  {
	bin := g.Cfg().GetString("backup.dumpbin")
	if len(bin) == 0 {
		bin = "mysqldump"
	}
	fmt.Println(gtime.New().Format("Y-m-d H:i:s") + "- dump database", db )
	file := gfile.Join(dir, db + ".sql")
	cmd := fmt.Sprintf("%s -u%s -p%s -h %s -P %s -B %s > %s", bin,g.Cfg().GetString("database.user"), g.Cfg().GetString("database.pass"), g.Cfg().GetString("database.host"), g.Cfg().GetString("database.port"), db, file)
	gproc.ShellExec(cmd)
}
