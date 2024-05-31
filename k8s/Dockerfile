FROM cimg/go:1.22.3

USER root
WORKDIR /usr/local/bin

# Install kubectl.
RUN curl -sSOL https://dl.k8s.io/release/v1.30.0/bin/linux/amd64/kubectl && \
  chmod +x kubectl

# Install helm.
ENV HELM_VERSION v3.14.4
RUN curl -sSOL https://get.helm.sh/helm-$HELM_VERSION-linux-amd64.tar.gz \
  && tar C . -xf helm-$HELM_VERSION-linux-amd64.tar.gz \
  && mv linux-amd64/helm .

# Install doctl.
ENV DOCTL_VERSION 1.107.0
RUN curl -sSOL https://github.com/digitalocean/doctl/releases/download/v$DOCTL_VERSION/doctl-$DOCTL_VERSION-linux-amd64.tar.gz \
  && tar C . -xf doctl-$DOCTL_VERSION-linux-amd64.tar.gz \
  && rm doctl-$DOCTL_VERSION-linux-amd64.tar.gz

USER circleci
WORKDIR /home/circleci/

# Install go tools.
RUN go install golang.stackrox.io/kube-linter/cmd/kube-linter@v0.6.8 && \
  go install github.com/zegl/kube-score/cmd/kube-score@v1.18.0 \
  && go clean --cache -testcache -fuzzcache -modcache

# Install ruby.
RUN curl -sSL https://rvm.io/mpapis.asc | gpg --import - && \
  curl -sSL https://rvm.io/pkuczynski.asc | gpg --import - && \
  curl -sSL https://get.rvm.io | bash -s stable && \
  bash -c "source $HOME/.rvm/scripts/rvm && rvm requirements && rvm install 3.3.1 && rvm use 3.3.1"
ENV PATH "/home/circleci/.rvm/rubies/ruby-3.3.1/bin:$PATH"

# Install gems.
RUN gem update bundler --force
