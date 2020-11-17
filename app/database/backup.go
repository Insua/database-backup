package database

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/gproc"
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
	fmt.Println("dump database ", db)
	file := gfile.Join(dir, db + ".sql")
	cmd := fmt.Sprintf("%s -u%s -p%s -h %s -P %s -B %s > %s", bin,g.Cfg().GetString("database.user"), g.Cfg().GetString("database.pass"), g.Cfg().GetString("database.host"), g.Cfg().GetString("database.port"), db, file)
	gproc.ShellExec(cmd)
	/*_, err := gproc.ShellExec(cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}*/
}
