FROM cimg/base:stable

USER root

# Install deps
RUN curl -s https://packagecloud.io/install/repositories/github/git-lfs/script.deb.sh | bash && \
    apt-get update && apt-get install --no-install-recommends -y \
    clang-format \
    git-lfs \
    && rm -rf /var/lib/apt/lists/*

# Install lfs
RUN git lfs install

WORKDIR /usr/local/bin

# Install dockerize.
ENV DOCKERIZE_VERSION v0.6.1
RUN curl -sSOL https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C . -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz

# Install hadolint.
RUN curl -sSOL https://github.com/hadolint/hadolint/releases/download/v2.10.0/hadolint-Linux-x86_64 && \
  mv hadolint-Linux-x86_64 hadolint && \
  chmod +x hadolint

# Install buf.
RUN curl -sSOL https://github.com/bufbuild/buf/releases/download/v1.6.0/buf-Linux-x86_64 && \
  mv buf-Linux-x86_64 buf && \
  chmod +x buf

# Install jq.
RUN curl -sSOL https://github.com/stedolan/jq/releases/download/jq-1.6/jq-linux64 && \
  mv jq-linux64 jq && \
  chmod +x jq

# Install yq.
RUN curl -sSOL https://github.com/mikefarah/yq/releases/download/4.26.1/yq_linux_amd64 && \
  mv yq_linux_amd64 yq && \
  chmod +x yq

USER circleci
WORKDIR /home/circleci/

# Install homebrew
RUN bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
ENV PATH "/home/linuxbrew/.linuxbrew/bin:$PATH"
RUN brew bundle --help
ENV LIBRARY_PATH "$LIBRARY_PATH:$(brew --prefix)/lib"

# Install xmake
RUN bash -c "$(curl -fsSL https://xmake.io/shget.text)"

# Install ruby.
RUN curl -sSL https://rvm.io/mpapis.asc | gpg --import - && \
  curl -sSL https://rvm.io/pkuczynski.asc | gpg --import - && \
  curl -sSL https://get.rvm.io | bash -s stable && \
  bash -c "source $HOME/.rvm/scripts/rvm && \
  rvm requirements && \
  rvm install 3.0.4 && rvm use 3.0.4"
ENV PATH "/home/circleci/.rvm/rubies/ruby-3.0.4/bin:$PATH"

# Install gems.
RUN gem update bundler --force
