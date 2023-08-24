#!/bin/sh
echo You probably rather want to install autodeployer2.
exit 10
BINARY=~/devel/admintools/install_autodeployer.sh
echo "Moved to admintools repository. executing ${BINARY} for your convenience"
exec ${BINARY} $@
