#!/bin/sh
addgroup deployers
CT=1
while [ $CT -le 80 ]; do
    echo $CT
    MU=deploy${CT}
    echo "User: ${MU}"
    useradd -s /bin/false -g deployers -M ${MU}
    usermod -g deployers -d /srv/autodeployer/deployments/${MU} ${MU}
    CT=$((CT+1))
done
