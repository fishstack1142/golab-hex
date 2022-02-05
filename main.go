package main

import (
	"bank/handler"
	"bank/repository"
	"bank/service"
	"fmt"
	"net/http"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux" //find out what is this
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func main() {
	initTimeZone()
	//use config
	initConfig()

	db := initDatabase()
	// dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
	// 	viper.GetString("db.username"),
	// 	viper.GetString("db.password"),
	// 	viper.GetString("db.host"),
	// 	viper.GetInt("db.port"),
	// 	viper.GetString("db.database"),
	// )

	// db, err := sqlx.Open(viper.GetString("db.driver"), dsn)
	// if err != nil {
	// 	panic(err)
	// }


	// db, err := sqlx.Open("mysql", "root:@tcp(localhost:3306)/banking")
	// if err != nil {
	// 	panic(err)
	// }

	customerRepository := repository.NewCustomerRepositoryDB(db) //db version


	// customerRepository := repository.NewCustomerRepositoryMock()
	customerService := service.NewCustomerService(customerRepository)
	customerHandler := handler.NewCustomerHandler(customerService)

	router := mux.NewRouter()

	router.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerId:[0-9]+}", customerHandler.GetCustomer).Methods(http.MethodGet)

	http.ListenAndServe(fmt.Sprintf(":%v", viper.GetInt("app.port")), router)

	// customers, err := customerService.GetCustomers()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(customers)

	// customers2, err := customerService.GetCustomer(1)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(customers2)


	// customerRepository := repository.NewCustomerRepositoryDB(db)
	
	// _ = customerRepository

	// customer, err := customerRepository.GetById(1)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(customer)

	// customer, err := customerRepository.GetAll()

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(customer)
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	time.Local = ict
}

func initDatabase() *sqlx.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.database"),
	)

	db, err := sqlx.Open(viper.GetString("db.driver"), dsn)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}