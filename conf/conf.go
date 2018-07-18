package conf

import (
	"encoding/json"
	"os"
)

type udp_srv_cnf struct {
	Addr string
}

type tcp_srv_cnf struct {
	Addr      string
	ReadLimit int
}

type booster_cnf struct {
	WorkerNum int
}

type Conf struct {
	AppID   string
	UUID    string
	UDP     *udp_srv_cnf
	TCP     *tcp_srv_cnf
	Booster *booster_cnf
}

func ReadConfig(confpath string) (*Conf, error) {
	file, _ := os.Open(confpath)
	decoder := json.NewDecoder(file)
	config := Conf{}
	err := decoder.Decode(&config)

	return &config, err
}
