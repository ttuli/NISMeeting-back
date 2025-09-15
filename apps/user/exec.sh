#!/bin/sh

error_exit() {
  echo "[ERROR] $1" >&2
  exit 1
}

SCRIPT_DIR=$(pwd)
mkdir -p ./log

RPC_NAME="userrpc"
API_NAME="userapi"

echo "[INFO] 开始编译 $RPC_NAME..."
cd ./rpc || error_exit "进入目录 ./rpc 失败"
go build -o ./build/$RPC_NAME || error_exit "编译 $RPC_NAME 失败"
./build/$RPC_NAME > ../log/$RPC_NAME.log 2>&1 &
pid=$!
if kill -0 $pid 2>/dev/null; then
  echo "[INFO] $RPC_NAME 已启动，日志输出在 ./log/$RPC_NAME.log (PID=$pid)"
else
  error_exit "运行 $RPC_NAME 失败"
fi

cd $SCRIPT_DIR

echo "[INFO] 开始编译 $API_NAME..."
cd ./api || error_exit "进入目录 ./api 失败"
go build -o ./build/$API_NAME || error_exit "编译 $API_NAME 失败"
./build/$API_NAME > ../log/$API_NAME.log 2>&1 &
pid=$!
if kill -0 $pid 2>/dev/null; then
  echo "[INFO] $API_NAME 已启动，日志输出在 ./log/$API_NAME.log (PID=$pid)"
else
  error_exit "运行 $API_NAME 失败"
fi

echo "[INFO] 所有服务启动成功 ✅"
