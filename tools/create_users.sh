#!/bin/sh
addgroup deployers
COUNT=$1
if [ "x$1" = "x" ]; then
    COUNT=80
fi

CT=1
while [ $CT -le ${COUNT} ]; do
    echo $CT
    MU=deploy${CT}
    echo "User: ${MU}"
    useradd -s /bin/sh -g deployers -M ${MU}
    usermod -g deployers -d /srv/autodeployer/deployments/${MU} ${MU}
    usermod -s /bin/sh ${MU}
    CT=$((CT+1))
done
