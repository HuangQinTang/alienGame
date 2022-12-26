基于ebitengine做的简单2d小游戏，飞机大战外星人

原教程：[一起用Go来做一个游戏](https://darjun.github.io/2022/11/15/godailylib/ebiten1/)

运行
- `go run .`
- main.exe，是已经编译好的window可执行文件。

浏览器运行
- `go test`，然后访问 http://127.0.0.1:8889/wasm_exec.html
- 利用wasm实现的，Go内置对wasm的支持，项目改动过后，需要重新编译wasm，`GOOS=js GOARCH=wasm go build -o main.wasm`。然后将位于$GOROOT/misc/wasm目录下的wasm_exec.html和wasm_exec.js文件拷贝到我们的项目目录下。注意wasm_exec.html文件中默认是加载名为test.wasm的文件，我们需要将加载文件改为main.wasm。

修改配置文件

- 配置文件config/config.json的修改需要用第三方包file2byteslice，重新打包二进制文件。`go install github.com/hajimehoshi/file2byteslice`
，然后在项目根目录执行`file2byteslice.exe -input ./config/config.json -output resources/config.go -package resources -var Config`
，生成的二进制文件在/resources下


