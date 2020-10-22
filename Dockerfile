FROM iits/vault-kubectl-helm:3.2.1-openapifix
WORKDIR /opt
COPY helm-charts-manager /opt