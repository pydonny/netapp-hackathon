package main

import (
 "os"
 "log"
 "github.com/BurntSushi/toml"

)


type Config struct { 
    Endpoint string
    bucket string
    AccessKey string
    SecretKey string
    ElsaticsearchHost string
    CameraCommand string
}


func ReadConfig(cf string) Config { 
    _, err := os.Stat(cf)
    if err != nil  { 
        log.Fatal("Config file is missing: ", cf) 
    }

    var config Config 
    if _, err := toml.DecodeFile(cf, &config); err != nil {
        log.Fatal(err)
    }

    return config
}

    
