package compareText 

import (
    "fmt"
    "log"
    "strings"
    "time"
    "os"
    "encoding/json"
    "github.com/yanyiwu/gojieba"
    "github.com/bregydoc/gtranslate"
)
const MaxLen int =  635

type CTReturn struct{
    WExist []string
    Num1    int
    Num2    int
    Percent float32
    WExist1 []string
    Percent1 float32
}

func getExMap() map[string]bool{
    exMap := make(map[string]bool,0)
    exMap["\""]=true
    exMap[" "]=true
    exMap[";"]=true
    exMap[","]=true
    exMap["，"]=true
    exMap["．"]=true
    exMap["."]=true
    exMap["、"]=true
    exMap["。"]=true
    exMap["《"]=true
    exMap["》"]=true
    return exMap
}

func compare2str(str1 string,str2 string) CTReturn{
    if str1 == "" || str2 == ""{
        return CTReturn{}
    }

    fmt.Println("==============新词识别==================")
    x := gojieba.NewJieba()
    defer x.Free()

    words := x.Cut(str1,true)
//    fmt.Println("新词识别:", strings.Join(words, "/"))
    wMap := make( map[string]bool,0)
    exMap := getExMap()
    for _,word := range words{
        if !exMap[word]{
    //        fmt.Println(word)
            wMap[word] = true
        }
    }

 //   fmt.Println(wMap)
    words1 := x.Cut(str2,true)
    wMap1 := make( map[string]bool,0)
    for _,word := range words1{
        if !exMap[word]{
     //       fmt.Println(word)
            wMap1[word] = true
        }
    }

  //  fmt.Println(wMap1)
    var wExist  []string
    for word,_ := range wMap1{
        if wMap[word]{
            fmt.Printf("word exist:%s\t",word)
            wExist = append(wExist,word)
        }
    }

    len1 := len(wMap)
    len2 := len(wMap1)
    percent := float32(len(wExist))/float32(len(wMap)) * 100

    var wExist1 []string
    for word,_ := range wMap{
        if wMap1[word]{
            wExist1 = append(wExist1,word)
        }
    }
    percent1 := float32(len(wExist1))/float32(len(wMap1)) * 100

    ret := CTReturn{WExist:wExist,Num1:len1,Num2:len2,Percent:percent,WExist1:wExist1,Percent1:percent1}        
    return ret
}

func Ct(str1 string,str2 string) string{
    ret := compare2str(str1,str2)
   // fmt.Printf("中文比对第二段话的词出现在第一段:%v,第一段字数:%d,第二段字数:%d,重复数:%d\n,重复率:%v\n",ret.WExist,ret.Num1,ret.Num2,len(ret.WExist),ret.Percent  )
    var resp []byte
    resp,err := json.Marshal(ret)
    if err != nil{
        log.Printf("json err:%s\n",err)
    }
    fmt.Printf("len1:%d,len2:%d\f",len(str1),len(str2))
    fmt.Printf("ct returning:%s\n",string(resp))
    return string(resp)
}

func CtEng(str1 string,str2 string) string{
    fmt.Println("call eng") 
    var str1T string
    var str2T string
    fmt.Printf("len1:%d,len2:%d\n",len(str1),len(str2))
    if len(str1)+len(str2) > MaxLen{
        fmt.Printf("length > 635 do translate twice\n")
       str1T = doTranslate(str1)
       str2T = doTranslate(str2)
    }else{
        fmt.Printf("length < 635 do translate once\n")
        cb := str1 + "\n||**||\n" + str2
        translated := doTranslate(cb)

        tStr := strings.Split(translated,"|| ** ||") 
        if len(tStr) < 2{
            tStr = strings.Split(translated,"||**||")
            if len(tStr) < 2{
                //save strings to file
                logErr(str1,str2)
                errRet := CTReturn{}
                resp,_ := json.Marshal(errRet)
                return string(resp) 
            } 
        } 
        str1T = tStr[0]
        str2T = tStr[1]
    }
    ret := compare2str(str1T,str2T)
    fmt.Printf("eng文比对第二段话的词出现在第一段:%v,第一段字数:%d,第二段字数:%d,重复数:%d\n,重复率:%v%",ret.WExist,ret.Num1,ret.Num2,len(ret.WExist),ret.Percent  )
    resp,err := json.Marshal(ret)
    if err != nil{
        log.Printf("json err:%s\n",err)
    }
    return string(resp) 
}

func logErr(str1,str2 string){
    t := time.Now()
    formatted := fmt.Sprintf("%d%02d%02d%02d%02d%02d",
    t.Year(), t.Month(), t.Day(),
    t.Hour(), t.Minute(), t.Second())
    fName := fmt.Sprintf("/tmp/ctErr%s.txt",formatted)
    f,_ := os.Create(fName)
    _,err := f.WriteString(fmt.Sprintf("str1:%s\n str2:%s",str1,str2))
    if err != nil{
        fmt.Println("write err log file failed")
    }


}

func doTranslate(str string) string{
    lStr := len(str)
    fmt.Printf("len:%d,maxLen:%d\n",lStr,MaxLen)
    if len(str) > MaxLen{
        runeStr := []rune(str)
        l := len(runeStr)/2
        fmt.Printf("half len:%d\n",l)
        str1 := string(runeStr[0:l])
        str2 := string(runeStr[l:])
        str1T := doTranslate(str1)
        str2T := doTranslate(str2)
        strAll := str1T + str2T
        return strAll
    }else{
        translated, err := gtranslate.TranslateWithParams(                                                                                                                                                                                   
            str,                                                                                                                                                                                                                              
            gtranslate.TranslationParams{                                                                                                                                                                                                    
                From: "zh-cn",                                                                                                                                                                                                               
                To:   "en",                                                                                                                                                                                                                  
            },                                                                                                                                                                                                                               
        )                                                                                                                                                                                                                                    
        if err != nil {                                                                                                                                                                                                                      
            // panic(err)                                                                                                                                                                                                                     
            fmt.Printf("error:%s\n",err.Error())                                                                                                                                                                                              
            logErr(str,"")
        }                                                                                                                                                                                                                                    

        //fmt.Println("translated:",translated)
        return translated

    }
}

