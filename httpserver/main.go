package main
import
(
    "fmt"
    "log"
    "net/http"
    "os"
    "strings"
)
//参考https://www.topgoer.com/%E7%BD%91%E7%BB%9C%E7%BC%96%E7%A8%8B/http%E7%BC%96%E7%A8%8B.html

func myHandle(w http.ResponseWriter, r *http.Request) {
  //w: 写给客户端的数据内容
  w.Write([]byte("This is a http web server."))

  //r：从客户端读到的内容
  fmt.Println("Header:",r.Header)
  fmt.Println("URL:",r.URL)
  fmt.Println("Method:",r.Method)
  fmt.Println("Host:",r.Host)
  fmt.Println("RemoteAddr:",r.RemoteAddr)
  fmt.Println("Body:",r.Body)

  //将request中的header写入response header
  for k,v:= range r.Header {
    fmt.Println(k,v)
    for _,vv := range v {
      w.Header().Set(k,vv)
    }
  }
 //将当前系统环境变量VERSION写入response header
 os.Setenv("VERSION", "0.0.0")
 version := os.Getenv("VERSION")
 w.Header().Set("VERSION",version)

 //Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
 //RemoteAddr在负载均衡时是LB的地址，需从其它字段获取真实IP
 log.Printf("Clent Real IP: %s, http return code:%d",getClientIP(r),http.StatusOK)
 //w.WriteHeader(http.StatusOK)

  //return ""
}

func healthz(w http.ResponseWriter, r *http.Request) {
  //w: 写给客户端的数据内容
  w.Write([]byte("Now health checking..."))
  w.WriteHeader(200)

}
func getClientIP(r *http.Request) string{
  ip := r.Header.Get("X-REAL-IP")
  if "" == ip {
    ip = strings.Split(r.RemoteAddr,":")[0]
  }

  return ip
}
func main () {
    fmt.Println("Hello wentingada!   \n\n   - from Golang .")
    //注册回调函数，该函数在客户端访问服务器时，自动被调用
    //http.HandleFunc("/", myHandle)
    mux := http.NewServeMux()
    mux.HandleFunc("/",myHandle)
    mux.HandleFunc("/healthz",healthz)

    //绑定服务器监听地址
    if err := http.ListenAndServe("127.0.0.1:3030", mux); err != nil {
      log.Fatalf("ListenAndServe Failed: %s\n", err.Error())
    }
}
