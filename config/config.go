package config

import (
	"encoding/json"
	"image/color"
	"log"
	"os"
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
}

func LoadConfig() *Config {
	f, err := os.Open("./config/config.json")
	if err != nil {
		log.Fatalf("os.Open failed: %v\n", err)
	}
	defer f.Close()

	var cfg Config
	err = json.NewDecoder(f).Decode(&cfg)
	if err != nil {
		log.Fatalf("json.Decode failed: %v\n", err)
	}

	return &cfg
}
