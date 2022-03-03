#!/bin/bash

# USAGE: ./build_kubernetes.sh <your_image> <your_namespace>
# e.g. ./build_kubernetes.sh docker-username/dashboard:test app-system

cd pkg/webui
npm install
ng build
cd ../..
release_dir=./dist/release
artifacts_dir=${release_dir}/artifacts
rm -r -f ${release_dir}
mkdir -p ${artifacts_dir}

GOOS="linux"
GOARCH="amd64"
platform="linux_amd64"
image="$1"
namespace="$2"

go_executable_output_file=${release_dir}/${platform}/dashboard

echo building Go executable for $GOOS $GOARCH, output will be $go_executable_output_file
env CGO_ENABLED=0 GOOS=$GOOS GOARCH=$GOARCH go build -a -o $go_executable_output_file ./cmd/areview/main.go

platform_release_dir=${release_dir}/${platform}

echo preparing Bhojpur Dashboard kubernetes release dir ${platform_release_dir}
mkdir -p ${platform_release_dir}/webui/
cp -r ./pkg/webui/dist ${platform_release_dir}/webui/
mv $go_executable_output_file ${platform_release_dir}

docker build -t ${image} .
docker push ${image}

kubectl delete -f ./deploy/test_dashboard.yaml -n ${namespace}

kubectl apply -f ./deploy/test_dashboard.yaml -n ${namespace}

sleep 5

appctl dashboard -k