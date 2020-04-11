#!/bin/sh

cd /robot/RoboMentor_SDK

export GOPATH=$PWD

cd /robot/RoboMentor_SDK/src/www

go build robot.go

echo "Build Success"