package users

import (
	"fmt"
	"time"
	"godesde0/modelos"
)

func AltaUsuario() {
	user := new(modelos.User)
	user.AddUser(10,"Pablo", time.Now(),true)
	fmt.Println(user)
}