package check_spam

import (
	"github.com/labstack/echo/v4"
	"math/rand"
	"net/http"
	"time"
)

var PullReasons = map[int]string{
	0: "blacklist",
	1: "stop_word",
	2: "obscene",
	3: "email",
	4: "link",
	5: "numeric",
	6: "quality_message",
}

var PullPermissions = map[int]string{
	0: "allowed",
	1: "not_allowed",
}

func (i *Implementation) GetPermissions(c echo.Context) error {
	req := new(GetPermissionsReq)
	err := c.Bind(req)
	if err != nil {
		return err
	}

	randSleep := rand.Intn(700)
	time.Sleep(time.Millisecond * time.Duration(randSleep))

	randPermission := PullPermissions[rand.Intn(10)%2]
	randReason := PullReasons[rand.Intn(100)%7]

	res := &GetPermissionsRes{
		Permission: randPermission,
	}

	if randPermission == "not_allowed" {
		res.Reason = randReason
	}

	return c.JSON(http.StatusOK, res)
}
