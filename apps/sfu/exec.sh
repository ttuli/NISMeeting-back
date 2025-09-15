#!/usr/bin/env bash

name="livekit-server"
pids=$(pgrep -f "$name")

if [ -z "$pids" ]; then
    ./livekit/bin/livekit-server --config ./livekit.yaml > ./log/livekit.log 2>&1 &
else 
    echo "[INFO] $name 已运行"
fi


