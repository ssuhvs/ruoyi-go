package lib_net

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
)

func GetClientRealIP(c *gin.Context) string {
	// 获取用户请求的 IP 地址
	ip := c.Request.Header.Get("X-Real-IP")
	if ip == "" {
		ip = c.Request.Header.Get("X-Forwarded-For")
	}
	if ip == "" {
		ip = c.Request.RemoteAddr
	}
	return ip
}

// 获取外网ip地址
func GetLocation(ip string) string {
	return ip
	//if ip == "127.0.0.1" || ip == "localhost" {
	//	return "内部IP"
	//}
	//resp, err := http.Get("https://restapi.amap.com/v3/ip?ip=" + ip + "&key=3fabc36c20379fbb9300c79b19d5d05e")
	//if err != nil {
	//	panic(err)
	//
	//}
	//defer resp.Body.Close()
	//s, err := ioutil.ReadAll(resp.Body)
	//fmt.Printf(string(s))
	//
	//m := make(map[string]string)
	//
	//err = json.Unmarshal(s, &m)
	//if err != nil {
	//	fmt.Println("Umarshal failed:", err)
	//}
	//if m["province"] == "" {
	//	return "未知位置"
	//}
	//return m["province"] + "-" + m["city"]
}

// 获取局域网ip地址
func GetLocaHost() string {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.Interfaces failed, err:", err.Error())
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()

			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						return ipnet.IP.String()
					}
				}
			}
		}

	}
	return ""
}
