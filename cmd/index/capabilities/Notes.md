# Notes

1. create secret with kubeconfig with name `kube-config-secret` in `openshift-operator-lifecycle-manager` namespace
2. create secret with dockerconfig to access private registries(registrey.connect.redhat.com, registry.redhat.io) with name `registry-redhat-dockerconfig` in `openshift-operator-lifecycle-manager` namespace
3. Add scc policy to default serviceaccount in `openshift-operator-lifecycle-manager` namespace by running:`oc adm policy add-scc-to-user privileged -z default -n openshift-operator-lifecycle-manager`