基于ebitengine做的简单2d小游戏，飞机大战外星人

原教程：[一起用Go来做一个游戏](https://darjun.github.io/2022/11/15/godailylib/ebiten1/)

运行 `go run .`

配置文件config/config.json的修改需要用第三方包file2byteslice，重新打包二进制文件。
- `go install github.com/hajimehoshi/file2byteslice`

在项目根目录执行
- `file2byteslice.exe -input ./config/config.json -output resources/config.go -package resources -var Config`

生成的二进制文件在/resources下
