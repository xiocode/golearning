#!/bin/bash
git pull remote master
git add .
git commit -m $1
git push remote

