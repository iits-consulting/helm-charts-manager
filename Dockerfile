FROM iits/vault-kubectl-helm:3.2.1-openapifix-stablerepofix
ARG VERSION
ARG GITHUB_REPOSITORY
WORKDIR /opt
ENV DOWNLOAD_LINK="https://github.com/${GITHUB_REPOSITORY}/releases/download/${VERSION}/helm-charts-manager_${VERSION}_linux_amd64.tar.gz"
RUN echo $DOWNLOAD_LINK
RUN curl -LO $DOWNLOAD_LINK
RUN tar -zxvf helm-charts-manager*
RUN mv helm-charts-manager* /usr/local/bin/