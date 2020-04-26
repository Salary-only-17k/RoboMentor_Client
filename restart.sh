#!/bin/sh

cd /robot/RoboMentor_Client

export GOPATH=$PWD

cd /robot/RoboMentor_Client/src/www

go build main.go

go build robot.go

sudo supervisorctl reload

sudo supervisorctl restart RoboMentorClient

sudo supervisorctl restart Robot

echo "Restart Success"