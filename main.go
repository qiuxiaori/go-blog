package main

import (
	"log"
	"net/http"
	"time"

	"github.com/qiuxiaori/go-blog/global"
	"github.com/qiuxiaori/go-blog/internal/routers"
	"github.com/qiuxiaori/go-blog/pkg/logger"
	"github.com/qiuxiaori/go-blog/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init config err %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
}

// @title 博客系统
// @version 1.0
// @description Go 语言编程之旅：一起用 Go 做项目
// @termsOfService https://github.com/go-programming-tour-book
func main() {
	// gin.SetMode(global.ServerSetting.RunMode)
	global.Logger.Infof("%s: go-programming-tour-book/%s", "eddycjy", "blog-service")

	r := routers.NewRouter()
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

	err = setting.ReadSetting("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	global.JWTSetting.Expire *= time.Second

	return nil
}

func setupLogger() error {

	print("this is setupLogger")
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  "logs/main.log",
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}
