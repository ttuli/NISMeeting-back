#!/usr/bin/env bash

name="livekit-server"
pids=$(pgrep -f "$name")

  if [ -z "$pids" ]; then
    echo "[INFO] $name 未运行"
  else
    echo "[INFO] 停止 $name (PID: $pids)"
    kill -9 $pids || echo "[ERROR] 停止 $name 失败"
  fi