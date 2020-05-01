#!/bin/sh

cd /robot/RoboMentor_Client

export GOPATH=$PWD

cd /robot/RoboMentor_Client/src/www

go build main.go

sudo supervisorctl reload

sudo supervisorctl restart RoboMentorClient

echo "Restart Success"