package main

import (
	"log"
	"net/http"
	"time"

	"github.com/qiuxiaori/go-blog/global"
	"github.com/qiuxiaori/go-blog/pkg/setting"
	"github.com/qiuxiaori/go-blog/src/router"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init config err %v", err)
	}
}

func main() {
	// gin.SetMode(global.ServerSetting.RunMode)

	r := router.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = setting.ReadSetting("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	return nil
}
