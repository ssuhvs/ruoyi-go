package lib_net

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"lostvip.com/utils/lib_logic"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

// ///////////////////////////////////////////////////////////////////
// 最好的代理方式：不要重新构建url 和 参数，
// 否则中文转码和参数存放位置都要进行匹配，比较麻烦
// ///////////////////////////////////////////////////////////////////
func ProxyWithUrlSame(c *gin.Context, host string) {
	fmt.Println("============= HttpProxyWithSameUrl start ==================")
	u := &url.URL{}
	u.Scheme = "http" //转发的协议，如果是https，写https.
	u.Host = host     //"127.0.0.1:7300"
	proxy := httputil.NewSingleHostReverseProxy(u)
	//重写出错回调
	proxy.ErrorHandler = func(rw http.ResponseWriter, req *http.Request, err error) {
		log.Printf(" http代理出错: %v", err)
		ret := fmt.Sprintf("http代理出错 %v", err)
		rw.Write([]byte(ret)) //写到body里
	}
	proxy.ModifyResponse = func(response *http.Response) error {
		response.Header.Set("proxy", "proxy by ssz")
		return nil
	}
	proxy.ServeHTTP(c.Writer, c.Request)
	fmt.Println("============= HttpProxyWithSameUrl over ==================")
}

// 比url不相同时，这种方式无效
//func ProxyHttp(c *gin.Context, host string, path string) {
//	fmt.Println("============= HttpProxyWithSameUrl start ==================")
//	url := &url.URL{}
//	url.Scheme = "http" //转发的协议，如果是https，写https.
//	url.Host = host     //"127.0.0.1:7300"
//	proxy := httputil.NewSingleHostReverseProxy(url)
//	//重写出错回调
//	proxy.ErrorHandler = func(rw http.ResponseWriter, req *http.Request, err error) {
//		log.Printf(" http代理出错: %v", err)
//		ret := fmt.Sprintf("http代理出错 %v", err)
//		rw.Write([]byte(ret)) //写到body里
//	}
//	proxy.ModifyResponse = func(response *http.Response) error {
//		response.Header.Set("proxy", "proxy by ssz")
//		return nil
//	}
//	c.Request.URL.Host = host
//	c.Request.URL.Path = path //修改旧的url为新的URL
//	proxy.ServeHTTP(c.Writer, c.Request)
//	fmt.Println("============= HttpProxyWithSameUrl over ==================")
//}

/**
 * 反向代理，转发到特定的URL上
 */
func ProxyWithUrlDifferent(c *gin.Context, targetUrl string, rawQuery string) {
	url, err := url.Parse(targetUrl)
	lib_logic.HasError1(err)
	lib_logic.Assert1(!strings.HasPrefix(targetUrl, "http"), "targetUrl 格式必须以http开头！")
	fmt.Println("==================rawQuery:" + rawQuery)
	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL = url
			req.URL.RawQuery = rawQuery
		},
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}
