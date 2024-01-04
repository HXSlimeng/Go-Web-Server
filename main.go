package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/HXSlimeng/Go-Web-Server/router"
)

func main() {
	r := router.SetupRouter()

   srv := &http.Server{
       Addr:    ":8080",
       Handler: r,
   }

   go func() {
       if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
           log.Fatalf("listen: %s\n", err)
       }
   }()

   // 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
   quit := make(chan os.Signal)
   signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
   <-quit
   log.Println("Shutdown Server ...")

   ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
   defer cancel()
   if err := srv.Shutdown(ctx); err != nil {
       log.Fatal("Server Shutdown:", err)
   }
   log.Println("Server exiting")
}
