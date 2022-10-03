#!/bin/bash

while getopts "n:" args; do
  case ${args} in
  n)
    name=${OPTARG}
    ;;
  *)
    exit
    ;;
  esac
done

full_path="${GOPATH}/src/playground/leetcode/${name}"

IFS="_"

# shellcheck disable=SC2162
read -a arr <<< "${name}"

md_title=''
go_module_name=''

arr_length=$((${#arr[@]}-1))
IFS=' ' # reset ifs

for i in "${!arr[@]}";
do
  if [ "${i}" -gt 0 ] && [ "${i}" -ne ${arr_length} ];then
    md_title+="${arr[i]} "
    go_module_name+="${arr[i]}_"
  elif [ "${i}" -eq ${arr_length} ];then
    md_title+="${arr[i]}"
    go_module_name+="${arr[i]}"
  fi
done

mkdir "${full_path}"

cd "${full_path}" || exit

touch "README.md"

{
  printf "# %s" "$md_title"
  printf "\n"
  printf "\n"
  printf "// TODO: add description"
  printf "\n"
} >>"README.md"

touch "main.go"

{
  printf "package main"
  printf "\n"
  printf "\n"
  printf "func main() {"
  printf "\n"
  printf "	// TODO: implement here"
  printf "\n"
  printf "}"
  printf "\n"
} >>"main.go"

touch "main_test.go"

{
  printf "package main"
} >>"main_test.go"

go mod init "${go_module_name}"

go get github.com/stretchr/testify
