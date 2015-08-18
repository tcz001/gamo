#!/bin/sh
set -e
# check to see if protobuf folder is empty
if [ ! -d "$HOME/protobuf/lib" ]; then
    git clone https://github.com/google/protobuf.git
    cd protobuf && ./autogen.sh && ./configure --prefix=$HOME/protobuf && make && make check && make install
else
    echo 'Using cached directory.';
fi
