#!/bin/bash
git pull git@github.com:xiocode/golearning.git master
git add .
git commit -m $1
git push git@github.com:xiocode/golearning.git 

