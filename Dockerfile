FROM iits/vault-kubectl-helm-limited-user:3.2.1-openapifix-stablerepofix
ARG VERSION
ARG GITHUB_REPOSITORY
WORKDIR /opt
USER root
ENV DOWNLOAD_LINK="https://github.com/${GITHUB_REPOSITORY}/releases/download/${VERSION}/helm-charts-manager_${VERSION}_linux_amd64.tar.gz"
RUN echo $DOWNLOAD_LINK && \
    curl -LO $DOWNLOAD_LINK && \
    tar -zxvf helm-charts-manager* && \
    mv helm-charts-manager* /usr/local/bin/
USER vault-kubectl-helm-user