#!/usr/bin/env bash


sed  -e 's/\n/""/g' -e 's/\[/\{/g' -e 's/\]/\}/g' -e 's/\},/\},/g' -e s#\"#\'#g  tmp.txt


# todo 根据文件生成文件夹，并将文件放入到中
### ll |grep -v 'test'  |grep .go |awk '{print $9}' |awk '{split($0,a,".go" ); print a[1]}' |xargs -I file sh -c 'mkdir file;mv file* file/;'