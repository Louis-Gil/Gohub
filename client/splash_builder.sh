#!/bin/bash

# 네이티브 스플래시 스크린 제거
echo "Removing existing native splash screen..."
dart run flutter_native_splash:remove

# 새로운 네이티브 스플래시 스크린 생성
echo "Creating new native splash screen from configuration file..."
dart run flutter_native_splash:create --path=./flutter_native_splash.yaml

echo "Splash screen update complete."