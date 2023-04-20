#!/bin/bash

while read p; do
    datasource_name="$(echo $p | cut -d'"' -f2)"
    pulumi_datasource_name="$(echo $datasource_name | sed 's/vcd_//' | sed -r 's/^(\w)/\U\1/' | sed -r 's/_(\w)/\U\1/g')"
    echo "\"$datasource_name\": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, \"get$pulumi_datasource_name\")},"
done <tf-datasource
