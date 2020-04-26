#!/bin/sh

cd /robot/RoboMentor_Client

export GOPATH=$PWD

cd /robot/RoboMentor_Client/src/www

go build robot.go

if [ $? -ne 0 ]; then
    echo "Build Error"
else
    echo "Build Success"
fi