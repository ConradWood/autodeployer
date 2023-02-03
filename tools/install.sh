#!/bin/sh
BINARY=~/devel/admintools/install_autodeployer.sh
echo "Moved to admintools repository. executing ${BINARY} for your convenience"
exec ${BINARY} $@
