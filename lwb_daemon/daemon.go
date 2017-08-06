package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	File, err := os.Create("log")
	if err != nil {
		fmt.Println("创建日志文件错误", err)
		return
	}
	log.SetOutput(File)
	ce("pid")
}
func ce(pid string) {
	File, err := os.OpenFile(pid, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Println(err)
		return
	}
	info, _ := File.Stat()
	if info.Size() != 0 {
		log.Println("pid file is exist")
		return
	}
	if os.Getppid() != 1 {
		log.Println(os.Args[0])
		args := append([]string{os.Args[0]}, os.Args[1:]...)
		os.StartProcess(os.Args[0], args, &os.ProcAttr{Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}})
		return
	}
	File.WriteString(fmt.Sprint(os.Getpid()))
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGUSR2)
	go HttpServer()
	for {
		s := <-c
		switch s {
		case syscall.SIGUSR2:
			fmt.Println("自定义型号.")
		case os.Interrupt:
			fmt.Println("安全退出")
			Exit(File)
		}
	}
}
func HttpServer() {
	http.HandleFunc("/", route)
	http.ListenAndServe(":1789", nil)
}
func route(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	fmt.Fprint(w, "Hello World\n")
}
func Exit(F *os.File) {
	F.Close()
	os.Remove(F.Name())
	fmt.Println("bye")
}
