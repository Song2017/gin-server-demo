function generate_swagger_lib() {
  input=$1
  output=$2
  type=$3
  pkg_name=$4
  pkg_version=$5

  # clean previous swagger_server
  echo "${green}Step1===${reset} remove old '$output'"
  rm -rf $output

  # generate server
  docker_image="openapitools/openapi-generator-cli:v5.0.0"

  validate_cmd="docker run -v ${PWD}:/local $docker_image  validate -i /local/$input"
  echo "RUN --- $validate_cmd"
  $validate_cmd

  gen_cmd="docker run -v ${PWD}:/local $docker_image generate \
    --input-spec /local/$input \
    --generator-name $type \
    --output /local/$output"
  gen_cmd="$gen_cmd -p apiPath=$pkg_name -p packageName=$pkg_name -p packageVersion=$pkg_version"

  echo "RUN --- $gen_cmd"
  $gen_cmd

  echo "GEN --- done"
}