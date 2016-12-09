package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os/user"
	"path"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

const (
	DefaultAuthSVC = "localhost:1502"
)

var ConfigDir string
var LogFile string
var CredentialsFile string
var ConfigFile string

type Config struct {
	Username string
	Password string
	AuthSVC  string
}

func init() {
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	home := u.HomeDir
	ConfigDir = path.Join(home, ".clawio")
	ConfigFile = path.Join(ConfigDir, "client.conf")
	LogFile = path.Join(ConfigDir, "client.log")
	CredentialsFile = path.Join(ConfigDir, "client.credentials")
}

func Get() *Config {
	c := &Config{}
	data, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		return c
	}
	if err := json.Unmarshal(data, c); err != nil {
		return c
	}
	return c
}

func Set(cfg *Config) {
	data, _ := json.MarshalIndent(cfg, "", "  ")
	ioutil.WriteFile(ConfigFile, data, 0600)
}

func GetToken() string {
	data, err := ioutil.ReadFile(CredentialsFile)
	if err != nil {
		return ""
	}
	return string(data)
}

func GetTokenContext() (context.Context) {
	md := metadata.Pairs("token", GetToken())
	return  metadata.NewContext(context.Background(), md)
}

func SetToken(token string) {
	if err := ioutil.WriteFile(CredentialsFile, []byte(token), 0600); err != nil {
		log.Fatalln(err)
	}
}
