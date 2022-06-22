#! /bin/bash
set -ex
cd `dirname $0`

mvn clean package
mkdir -p output/
cp target/function.jar output/
