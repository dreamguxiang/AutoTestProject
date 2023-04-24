package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Path string `json:"path"`
	Ip   string `json:"ip"`
	Port string `json:"port"`
}

func main() {
	con := loadConfig()
	//加载目录
	http.Handle("/", http.FileServer(http.Dir(con.Path)))

	log.Println("Server is running on " + con.Ip + ":" + con.Port)
	err := http.ListenAndServe(con.Ip+":"+con.Port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func isFileExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func loadConfig() Config {
	hasConfig := isFileExist("./config.json")
	if hasConfig == false {
		//生成默认配置文件
		con := Config{
			Path: "./",
			Ip:   "",
			Port: "8080",
		}
		newfile, err := os.Create("./config.json")
		if err != nil {
			log.Fatal(err)
		}
		defer newfile.Close()
		encoder := json.NewEncoder(newfile)
		//格式化
		encoder.SetIndent("", "    ")

		err = encoder.Encode(con)
		if err != nil {
			log.Fatal(err)
		}
	}
	file, err := os.Open("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}
