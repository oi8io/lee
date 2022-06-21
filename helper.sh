#!/usr/bin/env bash

str="[[0,1],[0,2],[3,5],[5,4],[4,3]]"
echo ${str}  | sed -e 's/\[/\{/g' -e 's/\]/\}/g'

