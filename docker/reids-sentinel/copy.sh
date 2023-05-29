#!/bin/bash

if [ ! $1 ]; then
    echo "Please input the ip of master"
    exit 1
fi

if [ ! -d config ]; then
    mkdir config
fi
if [ ! -d config/sentinel1 ]; then
    mkdir config/sentinel1
fi
if [ ! -d config/sentinel2 ]; then
    mkdir config/sentinel2
fi
if [ ! -d config/sentinel3 ]; then
    mkdir config/sentinel3
fi

sed -i '73c sentinel monitor mymaster '$1' 6379 2' sentinel.conf
cp sentinel.conf config/sentinel1/sentinel.conf
cp sentinel.conf config/sentinel2/sentinel.conf
cp sentinel.conf config/sentinel3/sentinel.conf