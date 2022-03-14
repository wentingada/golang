package main
import
(
    "fmt"
    "log"
    "net/http"
    "os"
    "strings"
    "os/signal"
    "syscall"
    "time"
    "io"

    "github.com/prometheus/client_golang/prometheus/promhttp"

    
)

func main () {
  fmt.Println("Hello wentingada!   %v\n\n ",time.Now())
  fmt.Println("Strtting up http server ...")
 //在程序开始，在注册表中注册延时metric指标。  
  metrics.Register()

  //注册回调函数，该函数在客户端访问服务器时，自动被调用
  mux := http.NewServeMux()
  mux.HandleFunc("/",myHandle)
  mux.HandleFunc("/healthz",healthz)
  //通过http暴露该指标
  mux.Handle("/metrics", promhttp.Handler())

  //引入srv对象，使得代码优雅关闭
  //close不会尝试关闭或等待websockects连接，不够优雅
  //优雅关闭使用shutdown
  srv := http.Server{
          Addr: ":80",
          Handler: mux,
          ReadTimeout:    10 * time.Second,
          WriteTimeout:   10 * time.Second,
  }
  

  //启动一个gorountine，绑定服务器监听地址
  go func() {
    log.Printf("Server start at 127.0.0.0:80")
    if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
      log.Fatalf("ListenAndServe Failed: %s\n", err.Error())
    }
}()
    //处理 SIGINT,SIGTERM,SIGKILL,SIGHUP,SIGQUIT信号
    ch := make(chan os.Signal)
    signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
    log.Println(<-ch)
    //优雅终止服务
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer func() {
      //终止前做点啥
      cancel()
    }()

    f err := srv.Shutdown(ctx); err != nil {
      log.Fatalf("Server Shutdown Failed:%+v", err)
    }
    //log.Println(s.Shutdown(ctx))
    //等待gorotine打印shutdown消息
    time.Sleep(time.Second * 10)
    log.Println("Server Exited Properly Done")

  }//main

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

  user := r.URL.Query().Get("user")
  //增加0-2秒的随机延时
  timer := metrics.NewTimer()
	defer timer.ObserveTotal()
  delay := randInt(10,2000)
	time.Sleep(time.Millisecond*time.Duration(delay))
	if user != "" {
		io.WriteString(w, fmt.Sprintf("hello [%s]\n", user))
	} else {
		io.WriteString(w, "hello [stranger]\n")
	}
	io.WriteString(w, "===================Details of the http request header:============\n")
  //将request中的header写入response header
  for k,v:= range r.Header {
    io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
    //fmt.Println(k,v)
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
  //w.Write([]byte("Now health checking..."))
  //w.WriteHeader(200)
  io.WriteString(w,"Health checking OK\n"

}
func getClientIP(r *http.Request) string{
  ip := r.Header.Get("X-REAL-IP")
  if "" == ip {
    ip = strings.Split(r.RemoteAddr,":")[0]
  }

  return ip
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}
