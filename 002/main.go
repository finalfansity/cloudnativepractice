package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// healthz
func Healthz(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "200\n")
}

// worker
func Worker(w http.ResponseWriter, req *http.Request)  {
	//获取request heaser 然后写入并返回
	os.Setenv("VERSION", "123")
	for k, v := range req.Header{
		s := ""
		for _, i := range v{
			s = s + i
		}
		fmt.Println(k,s)
		w.Header().Set(k, s)
	}
	// 获取系统env并写入response header
	env :=os.Getenv("VERSION")
	w.Header().Add("VERSION", env)
	//记录请求的ip和状态码
	w.WriteHeader(http.StatusOK)
	fmt.Printf("client host: %s, status code: %d\n", req.Host, http.StatusOK)
}

func main() {
	http.HandleFunc("/", Worker)
	http.HandleFunc("/healthz", Healthz)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}