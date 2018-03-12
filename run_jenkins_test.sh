#!/usr/bin/env bash

# 获取当前项目的目录情况
work_dir=${PWD##*/}
project_path=`pwd`
#echo ${work_dir}
#echo ${project_path}

# 将当前项目链到GOPATH
mkdir -p ~/go/src/
test_path=~/go/src/${work_dir}
# 如果GOPATH下存在该目录，则去pull代码
if [ ${test_path} ];then
    cd ${test_path}
    git pull
else
    # 将当前的项目链过去
    ln -s ${project_path} ${test_path}
fi

cd ${test_path}
go test ./...
