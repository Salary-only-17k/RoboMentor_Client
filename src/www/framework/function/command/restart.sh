#!/bin/sh

sudo supervisorctl restart RoboMentorClient

if [ $? -ne 0 ]; then
    echo "Restart Error"
else
    echo "Restart Success"
fi