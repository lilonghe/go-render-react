FROM golang:latest as build

WORKDIR /app
COPY . /app
RUN go mod download
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /app/render .
 
FROM alpine:3.14 as run
RUN apk add chromium-chromedriver

WORKDIR /app
COPY --from=build /app/ ./

# 复制中文字体到系统中，当作兜底字体
# COPY ./template/dist/zh.ttf /usr/share/fonts/TTF/

EXPOSE 8080

CMD [ "/app/render" ]
