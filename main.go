package main

import (
	"accountservice/dbclient"
	"accountservice/service"
	"flag"
	"fmt"

	"github.com/go-delve/delve/pkg/config"
	"github.com/spf13/viper"
)

var appName = "accountservice"

func init() {
	profile := flag.String("profile", "test", "Environment profile, something similiar to spring profile")
	configServerUrl := flag.String("configServerUrl", "http://configserver:8888", "Address to config server")
	configBranch := flag.String("configBranch", "master", "git branch to fetch configuration from")
	flag.Parse()

	viper.Set("profile", *profile)
	viper.Set("configServerUrl", *configServerUrl)
	viper.Set("configBranch", *configBranch)

}

func main() {
	fmt.Printf("Starting %v\n", appName)

	config.LoadConfigurationFromBranch(viper.GetString("configServerUrl"),
		appName,
		viper.GetString("profile"),
		viper.GetString("configBranch"))

	initializeBoltClient()
	service.StartWebServer(viper.GetString("server_port"))
}

func initializeBoltClient() {
	service.DBClient = &dbclient.BoltClient{}
	service.DBClient.OpenBoltDb()
	service.DBClient.Seed()
}
