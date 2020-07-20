#!/bin/bash

# -p = pull, -c = clone

#to pull: sh scripts/gitcomm.sh -p

#to clone: sh scripts/gitcomm.sh -c repotoclone

while getopts ":pc:" opt; do
  case $opt in
    p)
      eval `ssh-agent -s` > result.txt
      ssh-add /etc/ssh/id_rsa > result.txt
      git pull origin >&1
      ;;
    c)
      eval `ssh-agent -s`
      ssh-add /etc/ssh/id_rsa
      git clone $2
      ;;
    \?)
      echo "Invalid option: -$OPTARG" >&2
      exit 1
      ;;
    :)
      echo "Option -$OPTARG requires an argument." >&2
      exit 1
      ;;
  esac
done