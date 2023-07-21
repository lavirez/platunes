package conf

import (
    "encoding/json"
    "os"
    "fmt"
)


type Configuration struct{
    AMPQ_URL string `default:"amqp://rabbitmq:brabbit@rabbitmq-server:5672/"`
}

func InitiateConfiguration() Configuration { 
    conf := Configuration{}
    setDefault()
    return conf
}

func setDefault(conf Configuration) { 
    for i := 0; i < reflecr
}
