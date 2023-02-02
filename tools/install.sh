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

NEW_VERSION=`/tmp/autodeployer-server -ge_info print`
echo "Deploying version ${NEW_VERSION}"

ssh ${HOST} "sudo chown cnw /tmp/autodeployer-server  >/dev/null 2>&1 ; sudo chown cnw /usr/local/bin/autodeployer-server ; sudo rm /usr/local/bin/autodeployer-server"
rsync -pvraz --progress /tmp/autodeployer-server ${HOST}:/tmp/
autodeployer-client -server=${HOST} -stop
echo "installing autodeployer..."
ssh ${HOST} "sudo cp /tmp/autodeployer-server /usr/local/bin ; sudo chown cnw /usr/local/bin/autodeployer-server ; sudo chmod 755 /usr/local/bin/autodeployer-server ;  sudo systemctl restart autodeployer"
echo "done"

CUR_VERSION=`autodeployer-client -server=${HOST} -print_version`
while [ "${CUR_VERSION}" != "${NEW_VERSION}" ]; do
    echo "Waiting for ${CUR_VERSION} to be upgraded to ${NEW_VERSION}..."
    sleep 1
    CUR_VERSION=`autodeployer-client -server=${HOST} -print_version`
done

echo "Waiting for autodeployer to be listed in deploymonkey..."
CODE=1
while [ ${CODE} != 0 ]; do
    deploymonkey-client -list_deployers|grep -q ${HOST}
    CODE=$?
done

deploymonkey-client -apply_suggest
