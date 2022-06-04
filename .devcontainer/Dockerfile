# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.145.1/containers/go/.devcontainer/base.Dockerfile

# [Choice] Go version: 1, 1.15, 1.14
ARG VARIANT="1.18"
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

# [Option] Install Node.js
ARG INSTALL_NODE="true"
ARG NODE_VERSION="lts/*"
RUN if [ "${INSTALL_NODE}" = "true" ]; then su vscode -c "source /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi


ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update \
    && apt-get -y install --no-install-recommends apt-utils 2>&1
# Install git, process tools, lsb-release (common in install instructions for CLIs)
# Also dependencies for telepresence
RUN apt-get -y install git procps lsb-release sudo iptables curl

# Install kubectl
RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl
RUN chmod +x ./kubectl
RUN mv ./kubectl /usr/bin/kubectl

# Install Telepresence
RUN curl -fL https://app.getambassador.io/download/tel2/linux/amd64/latest/telepresence -o telepresence && \
    install -o root -g root -m 0755 telepresence /usr/local/bin/telepresence

# FIX error: failed to open TUN device /dev/net/tun: no such file or directory
# RUN mkdir -p /dev/net/ && mknod /dev/net/tun c 10 200 && chmod 600 /dev/net/tun

# Copy in kubectl config file and set environment variables
# RUN mkdir -p /root/.kube
# COPY kubeconfig /root/.kube/config



# ENV KUBECONFIG=/root/.kube/config

# Install shell & k8s tools
RUN apt install -y fish fonts-powerline sshfs
# RUN apt install -y kubectx fish fonts-powerline

# RUN kubectx "cell-001"

# RUN curl --silent -L http://get.oh-my.fish > /tmp/omf-install \
RUN curl --silent -L https://raw.githubusercontent.com/oh-my-fish/oh-my-fish/master/bin/install > /tmp/omf-install \
    && fish /tmp/omf-install --noninteractive --path=/usr/local/bin/omf --config=~/.config/omf \
    && rm /tmp/omf-install 

ENV SHELL /usr/bin/fish
RUN fish -c "omf install bobthefish"
RUN mkdir -p /root/.config/fish && echo 'set -gx KUBECONFIG (echo (find ~/.kube -type f -name config.\*.yaml) | sed 's/[[:space:]]/:/g')' > /root/.config/fish/config.fish

RUN git clone https://github.com/txn2/kubefwd.git \
    && cd ./kubefwd \
    && go build ./cmd/kubefwd/kubefwd.go \
    && go install ./cmd/kubefwd/kubefwd.go

RUN wget https://github.com/fullstorydev/grpcurl/archive/v1.8.5.tar.gz \
    && tar -zxvf v1.8.5.tar.gz \
    && cd grpcurl-1.8.5/cmd/grpcurl/ \
    && go build \
    && mv ./grpcurl /usr/bin/ \
    && cd ../../.. \
    && rm -R ./grpcurl-1.8.5

RUN apt install -y iputils-ping dnsutils

# Clean up
RUN apt-get autoremove -y \
    && apt-get clean -y \
    && rm -rf /var/lib/apt/lists/*
    
ENV DEBIAN_FRONTEND=dialog
