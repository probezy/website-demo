package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	hostname := r.Host
	publicIP := r.Header.Get("X-Real-IP")
	XForwardedFor := r.Header.Get("X-Forwarded-For")
	XForwardedProto := r.Header.Get("X-Forwarded-Proto")
	body := `<!DOCTYPE html>
		    <html>
        		<body>
                	<h1>这是一个测试站点8888</h1>
                	<p>We are testing cpolar reverse proxy server.</p>
				 	<a href="/template/test.iso">下载test.iso</a>
					<p>HOST: ` + hostname + `</p> 
					<p>X-Real-IP: ` + publicIP + `</p>
					<p>X-Forwarded-For: ` + XForwardedFor + `</p>
					<p>X_Forwarded_Proto: ` + XForwardedProto + `</p>
					<p>CODE:MAGIC8888</p>
				</body>
			</html>`
	w.WriteHeader(200)
	w.Write([]byte(body))
}

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func main() {

	var httpAddr string

	flag.StringVar(&httpAddr, "httpAddr", ":8888", "website address listen port")

	flag.Parse()
	flag.Usage()

	http.Handle("/template/", http.StripPrefix("/template/", http.FileServer(http.Dir("./template"))))
	http.HandleFunc("/", home)
	http.HandleFunc("/hello", hello)
	fmt.Printf("web start listen port %s \n", httpAddr)
	err := http.ListenAndServe(httpAddr, nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	fmt.Println("done.")
}
