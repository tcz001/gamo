#!/bin/sh
set -ex
git clone https://github.com/google/protobuf.git
cd protobuf && ./autogen.sh && ./configure && make && make check && make install
