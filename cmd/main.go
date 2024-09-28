package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/luungs/luungs-api/cmd/luungs"
	"github.com/luungs/luungs-api/config"
	"github.com/luungs/luungs-api/database"
)

func main() {
    db, err := database.NewMySQLStorage(mysql.Config{
        User:                   config.Envs.DBUser,
        Passwd:                 config.Envs.DBPassword,
        Addr:                   config.Envs.DBAddress,
        DBName:                 config.Envs.DBName,

        Net:                    "tcp",
        AllowNativePasswords:   true,
        ParseTime:              true,
    })
    if err != nil {
        log.Fatal(err)
    }

    server := luungs.NewAPIServer(":8080", db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}    
}
