#! /bin/bash

if [ -z $1 ]; then
	echo "[ERROR]: $0 src_name dst_name"
	exit;
fi

src_file=$1
dst_file=${src_file/\.*/}

g++ $src_file -o out/$dst_file

echo "[RUN]: ./out/$dst_file"
