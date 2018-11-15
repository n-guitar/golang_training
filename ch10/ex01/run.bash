#!/bin/bash 
go build ./main.go

echo "jpeg"
./main -i jpeg < ./gopher.jpg > ./gopher2.jpg


echo "gif"
./main -i gif < ./gopher.gif > ./gopher2.gif


echo "png"
./main -i png < ./gopher.png > ./gopher2.png
