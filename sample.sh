#!/bin/bash

javac -d java java/example/*.java

go run cli/main.go -archive java/example/Sample.class


go run cli/main.go -archive java/example/SampleInterface.class