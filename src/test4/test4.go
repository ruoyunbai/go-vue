package main

import (
    "io/ioutil"
    "log"
    "net/http"
)

func main() {
    // 在/后面加上 index ，来指定访问路径
    http.HandleFunc("/index", index)
    http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))
    http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./js"))))
    http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./img"))))

    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
    content, _ := ioutil.ReadFile("./index.html")
    w.Write([]byte(content))
}


