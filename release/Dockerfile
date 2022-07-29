FROM cimg/node:18.7.0

USER root

# Install deps
RUN curl -s https://packagecloud.io/install/repositories/github/git-lfs/script.deb.sh | bash && \
    apt-key adv --keyserver keyserver.ubuntu.com --recv-key C99B11DEB97541F0 && \
    apt-add-repository https://cli.github.com/packages && \
    apt-get update && apt-get install --no-install-recommends -y \
    software-properties-common \
    git-lfs \
    gh \
    && rm -rf /var/lib/apt/lists/*

# Setup deps
RUN npm i -g standard-version && \
    git lfs install

# Add release script
COPY release /usr/local/bin/
RUN chmod +x /usr/local/bin/release

USER circleci

# Disable prompt
RUN gh config set prompt disabled

# Setup git
RUN git config --global user.email "alexrfalkowski@gmail.com" && \
    git config --global user.name "Alejandro Falkowski"
