```
cd template
npm i && npm run build

cd ..
go build && ./example
```

## Docker
`docker build -t test . && docker run -p 8080:8080 test`  
View [http://127.0.0.1:8080/sample.pdf](http://127.0.0.1:8080/sample.pdf)

## Notice

### 驱动
运行环境需要安装 Chrome/Chromium 驱动
```
# alpine system
apk add chromium-chromedriver
```

### 非英文及数字文字渲染
运行环境需要安装中文字体包，否则中文字体会无法显示，解决方式有两种
1. 将文字拷贝到运行环境字体目录中（`cp ./template/dist/zh.ttf /usr/share/fonts/TTF/`）
2. 通过设置 CSS font-face 导入字体包

推荐两种方式同时使用，第一种作为兜底方案，因为 CSS 导入时必须将 dom 的 font-family 文字设为导入的字体，否则不会显示，因为不存在兜底字体。