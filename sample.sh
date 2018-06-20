#!/bin/bash

javac -d java java/example/*.java

cd java/
zip -r ../example.zip .
cd ..

go run cli/main.go -pa -pc -dbc all -archive example.zip
