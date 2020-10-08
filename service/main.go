package main

import (
    "net/http"
    "os"
    "strconv"
    "log"
    //"fmt"
    //"./compareText"
    "./htpServer"
)
//test : ab -n 20000 -c 200 "127.0.0.1:8000/inc?name=i"
func main() {

    //s := "《逍遥仙道》电影级国风手游，开启顶级视觉享受，画面精美绚丽，让游戏写意唯美；人物造型典雅时尚，真实动作，精致的战斗画面、华丽的技能特效、唯美的主线剧情，给您带来最完美的视觉盛宴，绝世神兵、萌宠坐骑、神兽觉醒等众多特色，满足你全部的想象，更有奇遇秘闻，体验情感羁绊，载酒人生，带你进入一个不一样的仙侠世界。"
    //s1 := "《逍遥江湖》是一款3D武侠MMORPG手机游戏，在这里，你可以顿悟绝世武学，自由搭配炫酷技能，夺取远古神级宝物，驯服超级洪荒异兽，碾压群英，甚至定义自己的爱恨情仇，肆意情缘。"
    //s1 := "《莽荒修仙录》是一款以修仙故事为背景的arpg手游，游戏酣畅淋漓的战斗操作，千变万化的装备搭配，带来不一样的修仙奇遇！御剑诛仙行万里，风花雪月总相宜。多种法宝显现、千段连击战斗、坐骑培养觉醒，助您遨游三界力战群仙！」"
    //tStr := strings.Split(translated,"||**||") 
    //words = x.CutAll(s)

    //ret := compareText.Ct(s,s1)
    //fmt.Printf("ret:%v",ret)
    //fmt.Println("===============Extract1 关键字=================")
    //keywords := x.ExtractWithWeight(s, 20)
    //fmt.Println("Extract 1:", keywords)
    //fmt.Println("===============Extract2 关键字=================")
    //keywords1 := x.ExtractWithWeight(s1, 20)
    //fmt.Println("Extract 2:", keywords1)
    server := htpServer.Server{htpServer.StartProcessManager(map[string]float32{"i":0,"j":0})}
    http.HandleFunc("/get", server.Get)
    http.HandleFunc("/set", server.Set)
    http.HandleFunc("/inc", server.Inc)
    http.HandleFunc("/ct", server.CT)

    portnum := 8881
    if len(os.Args) > 1 {
        portnum, _ = strconv.Atoi(os.Args[1])

    }
    log.Printf("compare text service Going to listen on port %d\n", portnum)
    log.Fatal(http.ListenAndServe(":"+strconv.Itoa(portnum), nil))
}


/*
    fmt.Println("===============tags=================")
    tags := x.Tag(s)
    fmt.Println("tags:", strings.Join(tags, "/"))
    fmt.Println("===============搜索引擎模式 不用词典:=================")
    fmt.Println("搜索引擎模式 不用词典:", strings.Join(se, "/"))
    se := x.CutForSearch(s,!use_hmm)
    fmt.Println("================搜索引擎模式 用词典:================")
    fmt.Println("搜索引擎模式 用词典:", strings.Join(seh, "/"))
    seh := x.CutForSearch(s,true)
    fmt.Println("================token搜索引擎模式================")
    wordinfos := x.Tokenize(s, gojieba.SearchMode, !use_hmm)
    fmt.Printf("token搜索引擎模式 :%v\n", wordinfos)
    fmt.Println("=================token默认模式===============")
    wordinfosD := x.Tokenize(s, gojieba.DefaultMode, !use_hmm)
    fmt.Printf("token默认模式 :%v\n", wordinfosD)
    fmt.Println("===============Extract 关键字=================")
    keywords := x.ExtractWithWeight(s, 10)
    fmt.Println("Extract:", keywords)
    fmt.Println("================================")
    */
