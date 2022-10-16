package config

const Host = "localhost"
const Port = ":8080"
const DatabaseHost = "localhost"
const DatabasePort = "5432"

var config map[string]string

func Init() {
	config = make(map[string]string)
	config["host"] = Host
	config["port"] = Port
	config["db_host"] = DatabaseHost
	config["db_port"] = DatabasePort
}

func GetConfig() map[string]string {
	return config
}
