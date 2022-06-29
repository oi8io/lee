#!/usr/bin/env bash

str="[[1,4,5],[1,3,4],[2,6]]"
echo ${str}  | sed -e 's/\[/\{/g' -e 's/\]/\}/g'

