package main

import (
    "log"
    "net/http"
    "strconv"
    "time"
    "io/ioutil"
    //"bytes"
    
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}




/*func filterRequest(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if strings.HasSuffix(r.URL.Path, "/") {
            http.NotFound(w, r)
            return
        }

        next.ServeHTTP(w, r)
    })
}*/



func main() {
    http.HandleFunc("/socket", func (w http.ResponseWriter, r *http.Request) {
        // Upgrade upgrades the HTTP server connection to the WebSocket protocol.
        conn, err := upgrader.Upgrade(w, r, nil)
        if err != nil {
            log.Fatal("upgrade failed: ", err)
        }
        defer conn.Close()

        // Continuosly read and write message
        for i := 0;; i++{
            /*mt, message, err := conn.ReadMessage()
            if err != nil {
                log.Println("read failed:", err)
                break
            }
            input := string(message)
            cmd := getCmd(input)
            msg := getMessage(input)
            if cmd == "add" {
                todoList = append(todoList, msg)
            } else if cmd == "done" {
                updateTodoList(msg)
            }
            output := "Current Todos: \n"
            for _, todo := range todoList {
                output += "\n - " + todo + "\n"
            }
            output += "\n----------------------------------------"
            message = []byte(output)*/
            log.Println(i);
            err = conn.WriteMessage(websocket.TextMessage, []byte(strconv.Itoa(i)))
            time.Sleep(1000 * time.Millisecond)
            if err != nil {
                log.Println("write failed:", err)
                break
            }
        }
    })
    
    http.HandleFunc("/api/post_buffer", func (w http.ResponseWriter, r *http.Request){
        body, err := ioutil.ReadAll(r.Body)
        if err != nil {
            log.Fatal(": ",);
        }
        log.Println("post_buffer", body)
    })
    
    http.HandleFunc("/api/post_string", func (w http.ResponseWriter, r *http.Request){
        body, err := ioutil.ReadAll(r.Body)
        if err != nil {
            log.Fatal(": ",);
        }
        log.Println("post_string", string(body))
    })
    
    
    
    http.HandleFunc("/api/get_buffer", func (w http.ResponseWriter, r *http.Request){
        log.Println("get_buffer")
        w.Write([]byte{3, 1, 4, 1, 5});
    })
    
    http.HandleFunc("/api/get_string", func (w http.ResponseWriter, r *http.Request){
        log.Println("get_string")
        w.Write([]byte("from server"));
    })
    
    
    
    http.Handle("/",http.FileServer(http.Dir("./static")));
    
    http.ListenAndServe(":9080", nil)
}














