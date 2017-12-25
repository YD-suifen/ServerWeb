package usersessionget

import (
	"github.com/astaxie/beego/context"
)


func UserGet(ctx *context.Context) string  {

	usersess := ctx.Input.CruSession.Get("Adminname")


	value, ok := usersess.(string)
	if ok {
		return value
	}

	return ""

}
