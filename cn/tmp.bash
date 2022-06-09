#!/usr/bin/env bash


sed  -e 's/\n/""/g' -e 's/\[/\{/g' -e 's/\]/\}/g' -e 's/\},/\},/g' -e s#\"#\'#g  tmp.txt