#!/bin/bash

if [ -z "$1" ]; then
    echo "Usage: $0 <tag_name>"
    exit 1
fi

TAG_NAME="$1"

# 태그가 로컬에 이미 존재하는 경우 삭제
git tag -d $TAG_NAME 2>/dev/null

# 원격 저장소의 태그 삭제
git push origin :refs/tags/$TAG_NAME

# 현재 커밋에 태그 설정
git tag $TAG_NAME

# 새로운 태그를 원격 저장소로 푸시
git push origin $TAG_NAME