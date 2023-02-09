package main

import (

	"fmt"
	"log"
	"time"

	// Echo Framewrokss
	"github.com/labstack/echo" 

	//Viper
	"github.com/spf13/viper"
    
	//Database
	"database/sql"
	"github.com/go-sql-driver/mysql"

	//Repository
	_taskRepo "To_Do_App/Task/repository"
	_userRepo "To_Do_App/User/repository"

	//Usecase
	_taskUsecase "To_Do_App/Task/usecase"
	_userUsecase "To_Do_App/User/usecase"

	// Delivery
	_taskHttpDelivery "To_Do_App/Task/delivery/http"
)

func init(){

	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}
}

func main(){
	log.Println(viper.GetString(`database.host`))
	
	cfg := mysql.Config{
        User:      viper.GetString(`database.user`),
        Passwd:    viper.GetString(`database.pass`),
        Net:       viper.GetString(`database.net`),
        Addr:      fmt.Sprintf("%s:%d",viper.GetString(`database.host`), viper.GetInt(`database.port`)),
        DBName:    viper.GetString(`database.name`),
		ParseTime: viper.GetBool(`database.parsetime`),
    }

    // Get a database handle.
    var err error
    dbConn, err := sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := dbConn.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")

	defer func(){

		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()


	// Echo Framewrok
	e := echo.New()

	// Repository
	taskRepo := _taskRepo.NewMysqlTaskRepo(dbConn)
	userRepo := _userRepo.NewMysqlUserRepo(dbConn)
	
	//Usecase 
	timeoutContext := 10 * time.Second
	taskUsecase := _taskUsecase.NewTaskUsecase(taskRepo, timeoutContext)
	userUsecase := _userUsecase.NewUserUsecase(userRepo, timeoutContext)

	//Delivery
	_taskHttpDelivery.NewTaskHandler(e,taskUsecase,userUsecase) 
	
	e.Logger.Fatal(e.Start(viper.GetString("server.address")))	

}