package database

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/gproc"
	"github.com/gogf/gf/os/gtime"
	"os"
)

func Backup(dir string, dbs []string)  {
	for _, v := range dbs {
		db(dir, v)
	}
}

func db(dir, db string)  {
	path := gfile.Join(dir, db)
	if err := gfile.Mkdir(path); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	d := g.DB()
	_, errUse := d.Exec("USE `"+db+"`")
	if errUse != nil {
		fmt.Println(errUse)
		os.Exit(1)
	}
	rows, errShow := d.Query("SHOW TABLES")
	if errShow != nil {
		fmt.Println(errShow)
		os.Exit(1)
	}
	fmt.Println(gtime.Now().Format("Y-m-d H:i:s") + "- dump database", db)
	var tableName string
	tables := make([]string, 0)
	for rows.Next() {
		if err := rows.Scan(&tableName); err == nil {
			tables = append(tables, tableName)
		}
	}
	for _, v := range tables {
		table(path, db, v)
	}
}

func table(dir, db, table string)  {
	bin := g.Cfg().GetString("backup.dumpbin")
	if len(bin) == 0 {
		bin = "mysqldump"
	}
	file := gfile.Join(dir, table + ".sql")
	cmd := fmt.Sprintf("%s --skip-opt -q -e --single-transaction -u%s -p%s -h %s -P %s -B %s --tables %s > %s", bin,g.Cfg().GetString("database.user"), g.Cfg().GetString("database.pass"), g.Cfg().GetString("database.host"), g.Cfg().GetString("database.port"), db, table, file)
	gproc.ShellExec(cmd)
}
