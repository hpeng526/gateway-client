#!/bin/sh
n2g=`iwinfo ra0 assoclis| grep ago | awk '{print $1}'`
n5g=`iwinfo rai0 assoclis| grep ago | awk '{print $1}'`

total=

for item in $n5g;
do
    if test -z "$total"
    then
        total=$item
    else
        total=$item"|"$total
    fi
done;

for item in $n2g;
do
    if test -z "$total"
    then
        total=$item
    else
        total=$item"|"$total
    fi
done;

c_d=`cat /tmp/dhcp.leases | awk '{print $4"#"$2}'|grep -i -E "$total"`

echo $c_d
