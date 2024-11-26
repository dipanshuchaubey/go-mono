package auth

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	xormadapter "github.com/casbin/xorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
)

func Auth() {
	// Initialize a Xorm adapter with MySQL database.
	a, err := xormadapter.NewAdapter("mysql", "root:root@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("error: adapter: %s", err)
	}

	m, err := model.NewModelFromFile("./model.conf")
	if err != nil {
		log.Fatalf("error: model: %s", err)
	}

	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		log.Fatalf("error: enforcer: %s", err)
	}

	data, _ := e.GetFilteredPolicy(0, "john")
	fmt.Println(data)

	sub := "john"
	obj := "data1"
	act := "read"

	ok, err := e.Enforce(sub, obj, act)
	if err != nil {
		log.Fatalf("error: enforce: %s", err)
	}

	if ok {
		fmt.Println("Allow the request")
	} else {
		fmt.Println("Deny the request")
	}
}
