package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
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
					<p>POST 提交表单测试</p>					
					<form method="post" action="/login">
    					<p>名字： <input type="text" name="username" value="xiaoMing" /></p>
    					<p>密码： <input type="password" name="password"/></p>
						<p>
							<input type="submit">
							<input type="reset">
						</p>
					</form>
				</body>
			</html>`
	w.WriteHeader(200)
	w.Write([]byte(body))
}

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello "+time.Now().String()+"\n")
}

// 结构体解析
func postHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // method: POST
	//param := &struct {
	//	Username string `json:"username"`
	//}{}
	fmt.Fprintln(w, r.FormValue("username"))
	//只解析form 中的参数 FormValue enctype="application/x-www-form-urlencoded"
	fmt.Fprintln(w, r.FormValue("password"))

	// 通过json解析器解析参数
	// &struct { Username string "json:\"username\"" }{Username:"xiaoming"}

	w.Write([]byte("ok"))
}

func main() {

	http.Handle("/template/", http.StripPrefix("/template/", http.FileServer(http.Dir("./template"))))
	http.HandleFunc("/", home)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/login", postHandler)

	err := http.ListenAndServe(":8888", nil) //设置监听的端口
	//err := http.ListenAndServeTLS(":443", "/Users/michael/.acme.sh/myapp.cpolar.net/fullchain.cer", "/Users/michael/.acme.sh/myapp.cpolar.net/myapp.cpolar.net.key", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	fmt.Println("done.")
}
