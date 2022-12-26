package config

import (
	"alienGame/resources"
	"bytes"
	"encoding/json"
	"image/color"
	"log"
)

type Config struct {
	ScreenWidth       int        `json:"screenWidth"`       //游戏窗口宽度
	ScreenHeight      int        `json:"screenHeight"`      //游戏窗口高度
	Title             string     `json:"title"`             //游戏标题
	BgColor           color.RGBA `json:"bgColor"`           //游戏背景色
	ShipSpeedFactor   float64    `json:"shipSpeedFactor"`   //飞船左右移动时的速度
	BulletWidth       int        `json:"BulletWidth"`       //子弹的宽度
	BulletHeight      int        `json:"BulletHeight"`      //子弹的高度
	BulletSpeedFactor float64    `json:"BulletSpeedFactor"` //子弹的速度
	BulletColor       color.RGBA `json:"BulletColor"`       //子弹的颜色
	MaxBulletNum      int        `json:"MaxBulletNum"`      //同时存在的最大子弹数
	BulletInterval    int64      `json:"BulletInterval"`    //子弹发射间隔(时间戳：毫秒)
	AlienSpeedFactor  float64    `json:"AlienSpeedFactor"`  //外星人速度
	AlienInterval     int64      `json:"AlienInterval"`     //外星人创建间隔（时间戳：毫秒）
	TitleFontSize     int        `json:"TitleFontSize"`
	FontSize          int        `json:"FontSize"`
	SmallFontSize     int        `json:"SmallFontSize"`
	FailNum           int        `json:"FailNum"`    //未消灭的外星人数，达到该数判断游戏失败
	SuccessNum        int        `json:"SuccessNum"` //成功消灭的外星人数，达到该数判定游戏通关
}

func LoadConfig() *Config {
	// 读取文件方式获取配置
	//f, err := os.Open("./config/config.json")
	//if err != nil {
	//	log.Fatalf("os.Open failed: %v\n", err)
	//}
	//defer f.Close()

	//var cfg Config
	//err = json.NewDecoder(f).Decode(&cfg)
	//if err != nil {
	//	log.Fatalf("json.Decode failed: %v\n", err)
	//}

	// 将config.json打包成成二进制文件引入
	var cfg Config
	err := json.NewDecoder(bytes.NewReader(resources.Config)).Decode(&cfg)
	if err != nil {
		log.Fatalf("json.Decode failed: %v\n", err)
	}

	return &cfg
}
