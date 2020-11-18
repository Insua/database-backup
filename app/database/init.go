package database

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"os"
)

func Init() (dir string, dbs []string) {
	db := g.DB()
	rows, err := db.Query("SHOW DATABASES")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var dbName string
	for rows.Next() {
		if err := rows.Scan(&dbName); err == nil {
			dbs = append(dbs, dbName)
		}
	}

	if len(dbs) == 0 {
		fmt.Println("no databases to backup")
		os.Exit(1)
	}

	fmt.Println(gtime.Now().Format("Y-m-d H:i:s") + "- " +gconv.String(len(dbs)) + " databases to backup")

	time := gtime.Now().Format("Y-m-d-H-i-s")

	storagePath := gfile.RealPath(g.Cfg().GetString("backup.storage"))
	path := gfile.Join(storagePath, time)

	if err := gfile.Mkdir(path); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}


	return path, dbs
}
