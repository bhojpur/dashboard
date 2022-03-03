#!/bin/bash

# USAGE: ./build_standalone.sh <your-platform>
# e.g. ./build_standalone.sh windows_amd64

platforms=("$1")
release_dir=./dist/release

cd pkg/webui
npm install
ng build
cd ../..

for platform in "${platforms[@]}"
do
  platform_split=(${platform//_/ })
  GOOS=${platform_split[0]}
  GOARCH=${platform_split[1]}
  
  go_executable_output_file=${release_dir}/${platform}/dashboard
  if [ $GOOS = "windows" ]; then
    go_executable_output_file+='.exe'
  fi

  echo building Go executable for $GOOS $GOARCH, output will be $go_executable_output_file
  env CGO_ENABLED=0 GOOS=$GOOS GOARCH=$GOARCH go build -a -o $go_executable_output_file ./cmd/areview/main.go

  platform_release_dir=${release_dir}/${platform}

  echo preparing Bhojpur Dashboard standalone release dir ${platform_release_dir}
  mkdir -p ${platform_release_dir}/webui/
  cp -r ./pkg/webui/dist ${platform_release_dir}/webui/
  cp $go_executable_output_file ${platform_release_dir}

  rm -rf $HOME/.bhojpur/bin/web
  rm $HOME/.bhojpur/bin/dashboard.exe

  cp -r ${platform_release_dir}/webui $HOME/.bhojpur/bin/web
  mv $go_executable_output_file $HOME/.bhojpur/bin/
done

appctl dashboard