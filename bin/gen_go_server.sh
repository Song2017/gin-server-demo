#!/usr/bin/env bash

# configuration
input="api-server.json"
output="swagger_server"
type="go-gin-server"
pkg_name="apiserver"
pkg_version="1.1.0"

source bin/gen_swagger_lib.sh
generate_swagger_lib $input $output $type $pkg_name $pkg_version

# server file
echo '
package main

import server "apiserver/v1/nomad-api"

func main() {
	router := server.InitRouter()

	router.Run(":8000")
}' > $output/main.go

sed -e 's|https://test.cn/|http://localhost:8000/|' $output/api/openapi.yaml > $output/api/openapi.yaml.tmp
if [ -f $output/api/openapi.yaml.tmp ]; then
  rm -rf $output/api/openapi.yaml
  mv -f $output/api/openapi.yaml.tmp $output/api/openapi.yaml
fi