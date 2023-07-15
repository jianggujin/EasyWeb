package server

import (
	"fmt"
	"github.com/jianggujin/EasyWeb/internal/config"
	"github.com/jianggujin/EasyWeb/internal/log"
	"net"
	"net/http"
	"strings"
)

func StartServer() {
	mergeMineType()
	for _, ip := range getIPs() {
		log.Infof("server address: http://%s:%d, static dir: %s", ip, config.Config.Server.Port, config.Config.Server.Static)
	}
	// 静态文件目录
	fs := http.FileServer(http.Dir(config.Config.Server.Static))

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Infof("%s %s %s", r.RemoteAddr, r.Method, r.RequestURI)
		path := r.URL.Path
		index := strings.LastIndex(path, ".")
		if index > 0 {
			suffix := path[index+1:]
			mineType, ok := MineType[suffix]
			if ok {
				w.Header().Set("Content-Type", mineType)
			}
		}
		fs.ServeHTTP(w, r)
	})
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Server.Port), mux)
	if err != nil {
		log.PanicError(err)
	}
}

func mergeMineType() {
	for k, v := range config.Config.Server.MineType {
		MineType[k] = v
		delete(config.Config.Server.MineType, k)
	}
}

func getIPs() (ips []string) {
	interfaceAddr, err := net.InterfaceAddrs()
	if err != nil {
		log.Warnf("fail to get net interface addrs: %v", err)
		return ips
	}

	for _, address := range interfaceAddr {
		ipNet, ok := address.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}
	return ips
}
