#!/usr/bin/env bash

set -x

#TESTNET=test-net-2-0-3
#TESTNET=test-net-2-1-0

if [ -z "$TESTNET" ]; then
    echo TESTNET var needs to be defined
    exit 1
fi

gcloud compute ssh $TESTNET-seed-0     --command 'sudo systemctl restart dispatch-disgo-node'

sleep 5
gcloud compute ssh $TESTNET-delegate-0 --command 'sudo systemctl restart dispatch-disgo-node'

sleep 5
gcloud compute ssh $TESTNET-delegate-1 --command 'sudo systemctl restart dispatch-disgo-node'

sleep 5
gcloud compute ssh $TESTNET-delegate-2 --command 'sudo systemctl restart dispatch-disgo-node'

sleep 5
gcloud compute ssh $TESTNET-delegate-3 --command 'sudo systemctl restart dispatch-disgo-node'

sleep 5
gcloud compute ssh $TESTNET-delegate-3 --command 'sudo systemctl restart dispatch-disgo-node'

echo DONE.