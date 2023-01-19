#!/bin/sh
HOST=$1
if [ -z "$HOST" ]; then
    echo "host required"
    exit 10
fi

if [ ! -f /tmp/autodeployer-server ]; then
    wget -O /tmp/autodeployer-server https://www.conradwood.net/builds/downloads/artefactid/8/version/latest/dist/linux/amd64/autodeployer-server || exit 10
    chmod 755 /tmp/autodeployer-server || exit 10
fi
ssh ${HOST} "sudo chown cnw /tmp/autodeployer-server  >/dev/null 2>&1 ; sudo chown cnw /usr/local/bin/autodeployer-server ; sudo rm /usr/local/bin/autodeployer-server"
rsync -pvraz --progress /tmp/autodeployer-server ${HOST}:/tmp/
autodeployer-client -server=${HOST} -stop
echo "installing autodeployer..."
ssh ${HOST} "sudo cp /tmp/autodeployer-server /usr/local/bin ; sudo chown cnw /usr/local/bin/autodeployer-server ; sudo chmod 755 /usr/local/bin/autodeployer-server ;  sudo systemctl restart autodeployer"
echo "done"
