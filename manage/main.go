package main

import (
	"fmt"
	"os"
	"os/signal"

	_ "github.com/go-sql-driver/mysql"
	"github.com/valyala/fasthttp"
	"github.com/xiaofengshuyu/vpn-manager/manage/common"
	"github.com/xiaofengshuyu/vpn-manager/manage/router"
)

func init() {
	// load config
}

func main() {

	var (
		errc = make(chan error)
	)
	go func() {
		common.Logger.Info("server listen 8086")
		err := fasthttp.ListenAndServe(":8086", router.VPNManageRouter)
		if err != nil {
			errc <- err
		}
	}()
	go func() {
		common.Logger.Info("server listen 8087")
		err := fasthttp.ListenAndServe(":8087", router.UserAccessRouter)
		if err != nil {
			errc <- err
		}
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, os.Kill)
		errc <- fmt.Errorf("%s", <-c)
	}()
	fmt.Println(<-errc)
}
