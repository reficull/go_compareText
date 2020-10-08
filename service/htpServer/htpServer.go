package htpServer

import (
    "fmt"
    "log"
    "net/http"
    //    "os"
    "strconv"
    "encoding/json"
    "../compareText"
)

type CommandType int

const (
    GetCommand = iota
    SetCommand
    IncCommand
    CTCommand
)

type Command struct {
    ty        CommandType
    name      string
    str1    string
    str2    string
    val       float32
    eng     string 
    replyChan chan string
}

type Res struct{
    err string
    num int
}
type Server struct {
    Cmds chan<- Command
}


func StartProcessManager(initvals map[string]float32) chan<- Command {
    counters := make(map[string]float32)
    for k, v := range initvals {
        counters[k] = v
    }
    cmds := make(chan Command)
    go func() {
        for cmd := range cmds {
            switch cmd.ty {
            case GetCommand:
                if val, ok := counters[cmd.name]; ok {
                    cmd.replyChan <- fmt.Sprintf("%s",val)
                } else {
                    cmd.replyChan <- string("-1")
                }
            case SetCommand:
                counters[cmd.name] = cmd.val
                cmd.replyChan <- fmt.Sprintf("%s",cmd.val)
            case IncCommand:
                if _, ok := counters[cmd.name]; ok {
                    counters[cmd.name]++
                    cmd.replyChan <- fmt.Sprintf("%s",counters[cmd.name])
                } else {
                    cmd.replyChan <- string(-1)
                }
            case CTCommand:
                var ret string
                fmt.Println("eng:",cmd.eng)
                if cmd.eng == "1" {
                    fmt.Println("using eng:",cmd.eng)
                    ret = compareText.CtEng(cmd.str1,cmd.str2)
                }else{
                    fmt.Println("eng:",cmd.eng)
                    ret = compareText.Ct(cmd.str1,cmd.str2)
                }
                //fmt.Printf("ct command logic  s1:%s,s2:%s\n ret:%s",cmd.str1,cmd.str2,ret)
                cmd.replyChan <- ret
            default:
                //log.Fatal("unknown command type", cmd.ty)
            }
        }
    }()
    return cmds
}

func makeRes(str string) string{
    response := &Res{}
    response.err= "str2 not found"
    response.num = -1
    ret ,_:= json.Marshal(response)
    return string(ret)

}

func (s *Server) CT(w http.ResponseWriter, req *http.Request) {
    fmt.Printf("get req:%v",req)
    name := req.URL.Query().Get("name")
    err := req.ParseForm()
    if err != nil {
        //panic(err)
    }
    v := req.Form
    fmt.Printf("form:%v\n",v)
    str1 := v.Get("str1")
    str2 := v.Get("str2")
    eng := v.Get("eng")
    if len(str2) == 0{
        fmt.Fprintf(w,  makeRes("str2 not set"))
        return 
    }
    if len(str1) == 0{
        fmt.Fprintf(w,  makeRes("str2 not set"))
        return 
    }
//    log.Printf("get str1 len:%d,str2 len:%d\n",len(str1),len(str2))
    log.Printf("eng:%v",eng)
    replyChan := make(chan string)
    s.Cmds <- Command{ty:CTCommand,name:name,str1:str1,str2:str2,eng:eng,replyChan:replyChan}
    reply := <- replyChan
    fmt.Fprintf(w,  reply)
}


func (s *Server) Get(w http.ResponseWriter, req *http.Request) {
    log.Printf("get %v", req)
    name := req.URL.Query().Get("name")
    replyChan := make(chan string)
    s.Cmds <- Command{ty:GetCommand,name:name,replyChan:replyChan}
    reply := <- replyChan 
    fmt.Fprintf(w, "%s: %d\n", name, reply)
}

func (s *Server) Set(w http.ResponseWriter, req *http.Request) {
    log.Printf("set %v", req)
    name := req.URL.Query().Get("name")
    val := req.URL.Query().Get("val")

    intval, err := strconv.Atoi(val)
    fval := float32(intval)
    if err != nil {
        fmt.Fprintf(w, "%s\n", err)

    } else {
        replyChan := make(chan string)
        s.Cmds <- Command{ty:SetCommand,name:name,val:fval,replyChan:replyChan}
        _ = <- replyChan
        fmt.Fprintf(w, "ok\n")
    }
}

func (s *Server) Inc(w http.ResponseWriter, req *http.Request) {
    log.Printf("inc %v", req)
    name := req.URL.Query().Get("name")

    replyChan := make(chan string)
    s.Cmds <- Command{ty:IncCommand,name:name,val:0,replyChan:replyChan}
    //reply := <-replyChan

    fmt.Fprintf(w, "ok\n")

}
/*
func main() {
    server := Server{startProcessManager(map[string]int{"i":0,"j":0})}
    http.HandleFunc("/get", server.get)
    http.HandleFunc("/set", server.set)
    http.HandleFunc("/inc", server.inc)
    http.HandleFunc("/ct", server.ct)

    portnum := 8000
    if len(os.Args) > 1 {
        portnum, _ = strconv.Atoi(os.Args[1])

    }
    log.Printf("Going to listen on port %d\n", portnum)
    log.Fatal(http.ListenAndServe("localhost:"+strconv.Itoa(portnum), nil))

}
*/
