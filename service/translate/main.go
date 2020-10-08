package main
import (
    "fmt" 
    "github.com/bregydoc/gtranslate"
)


func main(){
    text := "《逍遥仙道》电影级国风手游，开启顶级视觉享受，画面精美绚丽，让游戏写意唯美；人物造型典雅时尚，真实动作，精致的战斗画面、华丽的技能特效、唯美的主线剧情，给您带来最完美的视觉盛宴，绝世神兵、萌宠坐骑、神兽觉醒等众多特色，满足你全部的想象，更有奇遇秘闻，体验情感羁绊，载酒人生，带你进入一个不一样的仙侠世界。"
    translated, err := gtranslate.TranslateWithParams(
        text,
        gtranslate.TranslationParams{
            From: "zh-cn",
            To:   "en",
        },
    )
    if err != nil {
        panic(err)
    }

    fmt.Printf("en: %s | ja: %s \n", text, translated)
    // en: Hello World | ja: こんにちは世界
}

