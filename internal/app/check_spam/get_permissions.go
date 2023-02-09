package check_spam

import (
	"github.com/labstack/echo/v4"
	"math/rand"
	"net/http"
	"time"
)

func (i *Implementation) GetPermissions(c echo.Context) error {
	req := new(GetPermissionsReq)
	err := c.Bind(req)
	if err != nil {
		return err
	}

	sleepRand := rand.Intn(1000)
	time.Sleep(time.Millisecond * time.Duration(sleepRand))

	res := new(GetPermissionsRes)

	return c.JSON(http.StatusOK, res)
}
