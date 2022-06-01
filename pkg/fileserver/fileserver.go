package fileserver

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

func FileServer(dir string, port, globalSleep int) {
	if dir == "" {
		log.Fatalln("必须指定文件夹dir")
	}
	isExists, err := pathExists(dir)
	if err != nil {
		log.Fatalln(err)
	}
	if !isExists {
		log.Fatalf("文件夹[%s]不存在\n", dir)
	}

	handler := http.NewServeMux()

	// ?sleep=1 睡眠sleep秒后返回
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)

		q := r.URL.Query()
		sleep, _ := strconv.Atoi(q.Get("sleep"))
		if globalSleep > sleep {
			sleep = globalSleep
		}
		if sleep > 0 {
			time.Sleep(time.Duration(sleep) * time.Second)
		}
		w.Header().Set("Cache-Control", "no-store")
		http.ServeFile(w, r, dir+r.URL.Path)
	})

	addrs := ipAddrs()
	if globalSleep > 0 {
		log.Printf("全局sleep: %d秒\n", globalSleep)
	}
	log.Printf("文件夹: %s\n", dir)
	log.Printf("监听端口: %d\n", port)
	for _, addr := range addrs {
		log.Printf("浏览地址: http://%s:%d", addr, port)
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func ipAddrs() []string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil
	}

	result := make([]string, 0)
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok {
			if ipnet.IP.To4() != nil {
				result = append(result, ipnet.IP.String())
			}
		}
	}
	return result
}
