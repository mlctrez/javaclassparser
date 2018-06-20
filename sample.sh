#!/bin/bash

javac -d java java/example/*.java

go run cli/main.go -archive java/example/Sample.class


go run cli/main.go -archive /Users/mattman/.m2/repository/org/elasticsearch/elasticsearch/1.3.2/elasticsearch-1.3.2.jar -pc | wc -l


# go run cli/main.go -archive java/example/SampleInterface.class