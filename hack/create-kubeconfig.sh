#!/usr/bin/env bash

# Copyright 2020 Authors of Arktos.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

LOG_DIR=${LOG_DIR:-"/tmp"}
KUBE_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
CERT_DIR=${CERT_DIR:-"/var/run/kubernetes"}
CERT_ROOTNAME=${CERT_ROOTNAME:-"host"}

source "${KUBE_ROOT}/hack/lib/util.sh"

# Ensure CERT_DIR is created for crt/key and kubeconfig
mkdir -p "${CERT_DIR}" &>/dev/null || sudo mkdir -p "${CERT_DIR}"
CONTROLPLANE_SUDO=$(test -w "${CERT_DIR}" || echo "sudo -E")
# There are following commands to run the script
# hack/create-kubeconfig.sh is to create a kubeconfig giving the following params
# hack/create-kubeconfig.sh extract key is to extract content from an existing kubeconfig.
if [ $# -lt 1 ] ; then
  echo "Not enough params"
else
  echo "The current operation is $1"
  case $1 in
    extract)
      if [ -n "$2" ]; then
        ${CONTROLPLANE_SUDO} grep $2 ${CERT_DIR}/admin.kubeconfig | awk '{print $2}'
      fi
      ;;
    extractall)
      ${CONTROLPLANE_SUDO} grep server ${CERT_DIR}/admin.kubeconfig | awk '{print $2}'
      ${CONTROLPLANE_SUDO} grep certificate-authority-data ${CERT_DIR}/admin.kubeconfig | awk '{print $2}'
      ${CONTROLPLANE_SUDO} grep client-certificate-data ${CERT_DIR}/admin.kubeconfig | awk '{print $2}'
      ${CONTROLPLANE_SUDO} grep client-key-data ${CERT_DIR}/admin.kubeconfig | awk '{print $2}'
      ;;
    create)
       ${CONTROLPLANE_SUDO} cat <<EOF > ${CERT_DIR}/${CERT_ROOTNAME}.kubeconfig
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: $3  
    server: $2
  name: local-up-cluster
contexts:
- context:
    cluster: local-up-cluster
    user: local-up-cluster
  name: local-up-cluster
current-context: local-up-cluster
kind: Config
preferences: {}
users:
- name: local-up-cluster
  user:
    client-certificate-data: $4
    client-key-data: $5
EOF
       ;;
    *)
      echo "Unknown operation"
      ;;
  esac
fi