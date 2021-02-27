#!/usr/bin/env sh

#start client
nc localhost 8081

#start server
#nc -l -p 8081 | awk '{print toupper($0)}'
#nc -l -p 8081 | awk '{print toupper($0) "some more bytes"}'


#udp client
#nc –u localhost 8888


#sudo hping3 --udp 192.168.1.193 -p 8081 -E test.tmp -d 6 -s 8082 -c 1
#nc -l -u 8082