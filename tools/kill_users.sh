#!/bin/sh
CT=1
while [ $CT -le 80 ]; do
    echo $CT
    MU=deploy${CT}
    echo "User: ${MU}"
    slay -9 ${MU}
    CT=$((CT+1))
done
