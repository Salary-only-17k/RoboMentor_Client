#!/bin/sh

cd /robot/RoboMentor_SDK

export GOPATH=$PWD

cd /robot/RoboMentor_SDK/src/www

go build main.go

sudo supervisorctl restart RoboMentorSDK

echo "MentorsService Restart Success"