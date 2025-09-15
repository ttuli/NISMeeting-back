#!/bin/sh

error_exit() {
  echo "[ERROR] $1" >&2
  exit 1
}

APP_NAME="fileapi"

go build -o ./build/$APP_NAME || error_exit "编译 $APP_NAME 失败"
./build/$APP_NAME > ../log$APP_NAME.log 2>&1 &
pid=$!
if kill -0 $pid 2>/dev/null; then
    echo "[INFO] $APP_NAME 已启动，日志输出在 ../log/$APP_NAME.log (PID=$pid)"
else 
    error_exit "运行 $APP_NAME 失败"
fi
