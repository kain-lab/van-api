package token

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"os"
	"testing"
	"van-api/bootstrap"
)

var tokenString string
var err error

func TestMain(m *testing.M) {
	os.Chdir("../..")
	cfg, err := bootstrap.LoadConfiguration()
	if err != nil {
		log.Fatalln(err)
	}
	Key = []byte(cfg.App.Key)
	Options = cfg.Token
	os.Exit(m.Run())
}

func TestMake(t *testing.T) {
	tokenString, err = Make("system", jwt.MapClaims{
		"username": "kain",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tokenString)
}

func TestVerify(t *testing.T) {
	result, claims, err := Verify("system", tokenString)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
	t.Log(claims)
}