package api

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"net/http"
	"time"
)

var srv http.Handler

func init() {
	e := echo.New()
	e.GET("/books", Test)
	e.GET("/aprs_passcode", Test)
	srv = e
}

func Test(c echo.Context) error {
	defer func() {
		if r := recover(); r != nil {
			c.Logger().Info("client resp error", r)
			c.String(200, "data error, please contact BH4FWA use Wechat.")
			return
		}
	}()
	str := "now time:" + time.Now().Format(time.DateTime)
	log.Infof("now time: %s", str)
	return c.String(200, str)
}

func MainFunc(w http.ResponseWriter, r *http.Request) {
	go func() {
		for {
			log.Info("now date:", time.Now().String())
			time.Sleep(time.Second * 5)
		}
	}()
	srv.ServeHTTP(w, r)
}
