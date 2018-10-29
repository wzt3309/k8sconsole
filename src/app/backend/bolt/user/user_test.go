package user

import (
	"github.com/boltdb/bolt"
	"github.com/stretchr/testify/assert"
	api "github.com/wzt3309/k8sconsole/src/app/backend/api"
	"log"
	"os"
	"path"
	"testing"
)

const dbfile = "tmp.db"

var db *bolt.DB

func init() {
	root := path.Join(os.Getenv("GOPATH"), "src/github.com/wzt3309/k8sconsole")
	tmp := path.Join(root, ".tmp/db")

	_, err := os.Stat(tmp)
	if os.IsNotExist(err) {
		err := os.MkdirAll(tmp, 0755)
		if err != nil {
			log.Fatalf("mkdir %s failed, err: %v", tmp, err)
		}
	}
	tmpdb := path.Join(tmp, dbfile)

	db, err = bolt.Open(tmpdb, 0600, nil)
	if err != nil {
		log.Fatalf("Cannot open db %s, err :%v", tmpdb, err)
	}
}

func TestNewService(t *testing.T) {
	_, err := NewUserService(db)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
}

func TestService_CreateUser(t *testing.T) {
	as := assert.New(t)

	svc, err := NewUserService(db)
	if err != nil {
		t.Fatal(err)
	}

	expected := &api.User{
		Username: "test01",
		Password: "123456",
		Role:     api.UserRole(0),
	}
	err = svc.CreateUser(expected)
	if err != nil {
		t.Fatal(err)
	}

	actual, err := svc.User(expected.ID)
	if err != nil {
		t.Fatal(err)
	}

	//t.Log(actual)

	as.Equal(expected, actual, "Can't create user")
}

func TestService_UpdateUser(t *testing.T) {
	as := assert.New(t)

	svc, err := NewUserService(db)
	if err != nil {
		t.Fatal(err)
	}

	user := &api.User{
		Username: "test01",
		Password: "123456",
		Role:     api.UserRole(0),
	}
	err = svc.CreateUser(user)
	if err != nil {
		t.Fatal(err)
	}

	expected := "777"
	user.Password = expected
	svc.UpdateUser(user.ID, user)

	actual, err := svc.User(user.ID)
	if err != nil {
		t.Fatal(err)
	}

	as.Equal(expected, actual.Password, "Can't update user")
}
