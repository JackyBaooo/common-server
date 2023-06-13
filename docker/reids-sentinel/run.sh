#!/bin/bash

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

sed -i '73c sentinel monitor mymaster 10.100.97.80 6380 2' sentinel.conf
cp sentinel.conf config/sentinel1/sentinel.conf
sed -i '10c port 26379' config/sentinel1/sentinel.conf
cp sentinel.conf config/sentinel2/sentinel.conf
sed -i '10c port 26380' config/sentinel2/sentinel.conf
cp sentinel.conf config/sentinel3/sentinel.conf
sed -i '10c port 26381' config/sentinel3/sentinel.conf