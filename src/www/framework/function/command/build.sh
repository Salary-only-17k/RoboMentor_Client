#!/bin/sh

cd /robot/RoboMentor_SDK

export GOPATH=$PWD

cd /robot/RoboMentor_SDK/src/www

go build robot.go

if [ $? -ne 0 ]; then
    echo "Build Error"
else
    echo "Build Success"
fi