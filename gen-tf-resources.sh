#!/bin/bash

while read p; do
    resource_name="$(echo $p | cut -d'"' -f2)"
    pulumi_resource_name="$(echo $resource_name | sed 's/vcd_//' | sed -r 's/^(\w)/\U\1/' | sed -r 's/_(\w)/\U\1/g')"
    echo "\"$resource_name\": {Tok: tfbridge.MakeResource(mainPkg, mainMod, \"$pulumi_resource_name\")},"

done <tf-resource
