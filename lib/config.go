package lib

import (
	"fmt"
	util "jetsend_opens/shared/helpers"
	"jetsend_opens/shared/log"
	"os"
)

func ParseConfiguration(configPath, logPath string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Execption ParseConfiguration err: ", err)
		}
	}()

	jmtafilePath := configPath + "/kafka.yaml"
	var kafka KafkaConfiguration
	util.ParseConfiguration(jmtafilePath, &kafka)

	servicefilePath := configPath + "/service.yaml"
	var service ServiceConfiguration
	util.ParseConfiguration(servicefilePath, &service)

	env := os.Getenv("JETSEND_ENV")
	ServiceConfig = Configuration{
		Kafka:   kafka.Env[env],
		Service: service.Env[env],
	}

	log.ProjectID = ServiceConfig.Service.ProjectID
	log.LogName = ServiceConfig.Service.LogName

	fmt.Println(env, " kafka: ", kafka.Env[env])
	fmt.Println(env, " service: ", service.Env[env])

	log.InitializeLogger(logPath)

}
