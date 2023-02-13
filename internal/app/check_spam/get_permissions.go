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

	res, _ := i.getPermissions(req)

	return c.JSON(http.StatusOK, res)
}

func (i *Implementation) getPermissions(req *GetPermissionsReq) (*GetPermissionsRes, error) {
	// искусственная нагрузка

	perNumber := 0
	for _, s := range req.MessageText {
		for j := 0; j < 10000; j++ {
			for k := 0; k < 1000; k++ {
				perNumber += int(s) + j + k
			}
		}
	}

	randSleep := rand.Intn(100)
	time.Sleep(time.Millisecond * time.Duration(randSleep))

	permission := PullPermissions[perNumber%2]
	res := &GetPermissionsRes{
		Permission: permission,
	}

	randReason := PullReasons[rand.Intn(100)%7]

	if permission == "not_allowed" {
		res.Reason = randReason
	}

	return res, nil
}
