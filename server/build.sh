#!/bin/bash

echo "Cleaning"
rm ./wildrain

echo "Building"
go build com.wildrain/wildrain

echo "Ricing."
rice -i com.wildrain/wildrain append  --exec wildrain

echo "Running"
./wildrain
