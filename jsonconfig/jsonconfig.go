package jsonconfig

import (
	"io"
	"io/ioutil"
	"os"
	"time"
	"encoding/json"
)

// Parser must implement ParseJSON
type Parser interface {
	ParseJSON([]byte) error
}

// Parser must implement ParseJSON
type Server struct {
	Hostname     string        `json:"hostname"`
	UseHTTP      bool          `json:"usehttp"`
	UseHTTPS     bool          `json:"usehttps"`
	HTTPPort     int           `json:"httpport"`
	HTTPSPort    int           `json:"httpsport"`
	CertFile     string        `json:"certfile"`
	KeyFile      string        `json:"keyfile"`
	ReadTimeout  time.Duration `json:"readtimeout"`
	WriteTimeout time.Duration `json:"writetimeout"`
}

type Database struct{
	Driver string `json:"driver"`
	Host 	 string `json:"host"`
	Port	 string	`json:"port"`
	DB		 string `json:"db"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Client struct {
	Port   int    `json:"port"`
	Prefix string `json:"prefix"`
}

type Backend struct {
	Port   int    `json:"port"`
	Prefix string `json:"prefix"`
}

// configuration contains the application settings
type Configuration struct {
	Session  interface{} 		 `json:"Session"`
	Database Database  	 		 `json:"Database"`
	Server   Server          `json:"Server"`
	Client   Client          `json:"Client"`
	Backend  Backend         `json:"Backend"`
	Env      string          `json:"Env"`
	BaseUrl  string          `json:"BaseUrl"`
	WhitelistIP []string		 `json:"WhitelistIP"`
}
// ParseJSON unmarshals bytes to structs
func (c *Configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}
// Load the JSON config file

func Load(configFile string, p Parser) error {
	var err error
	var input = io.ReadCloser(os.Stdin)
	if input, err = os.Open(configFile); err != nil {
		//log.Fatalln(err)
    return err
	}
	// Read the config file
	jsonBytes, err := ioutil.ReadAll(input)
	input.Close()
	if err != nil {
		//log.Fatalln(err)
    return err
	}

	// Parse the config
	if err := p.ParseJSON(jsonBytes); err != nil {
		//log.Fatalln("Could not parse %q: %v", configFile, err)
    return err
	}
  return nil
}
