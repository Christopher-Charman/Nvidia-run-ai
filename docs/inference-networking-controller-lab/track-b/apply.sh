#!/usr/bin/env bash
# Substitute ${DYNAMO_VERSION} (and ${EPP_IMAGE_TAG}) into a Track B manifest and apply it.
#   DYNAMO_VERSION=1.4.2 ./apply.sh 01-mocker-agg.yaml
set -euo pipefail
: "${DYNAMO_VERSION:?set DYNAMO_VERSION to the operator image tag from 00-check-dynamo.sh}"
export DYNAMO_VERSION
export EPP_IMAGE_TAG="${EPP_IMAGE_TAG:-1.4.0}"   # legacy Go EPP image tag used by Dynamo's own v1.4.2 GAIE example
for f in "$@"; do
  echo "== apply $f  (DYNAMO_VERSION=$DYNAMO_VERSION EPP_IMAGE_TAG=$EPP_IMAGE_TAG)"
  sed -e "s|\${DYNAMO_VERSION}|$DYNAMO_VERSION|g" -e "s|\${EPP_IMAGE_TAG}|$EPP_IMAGE_TAG|g" "$f" | kubectl apply -f -
done
