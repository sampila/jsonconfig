package jsonconfig

import(
  "testing"
  "os"
  "github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M){
  os.Exit(m.Run())
}

func TestLoadFileNotFound(t *testing.T){
  var config Configuration
  err := Load("/config/config.json", &config)
  assert.NotNil(t,err)
  assert.EqualValues(t, "open /config/config.json: no such file or directory", err.Error())
}

func TestReadFail(t *testing.T){
  var config Configuration
  err := Load("../config/config-empty.json", &config)
  assert.NotNil(t,err)
  assert.EqualValues(t, "unexpected end of JSON input", err.Error())
}

func TestParseError(t *testing.T){
  var config Configuration
  err := Load("../config/config-invalid.json", &config)
  assert.NotNil(t,err)
  assert.EqualValues(t, "invalid character '}' after top-level value", err.Error())
}

func TestValidFile(t *testing.T){
  var config Configuration
  err := Load("../config/config-valid.json", &config)
  assert.Nil(t,err)
  //assert.EqualValues(t, `{map[Name:commarcesess Options:map[Domain: HttpOnly:true MaxAge:28800 Path:/ Secure:false] SecretKey:@C0mm4rce000] map[Dbdriver:mysql Dbhost:localhost Dbname:golang_try Password:aptikma Username:root] {localhost true false 8081 443 keys/tls/server.crt keys/tls/server.key 10ns 10ns} {4000 /} {8081 /api} development http://localhost:8081}`, string(config))
}
