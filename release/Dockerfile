FROM circleci/node:15-buster

USER root

# Install software-properties-common
RUN apt-get update && apt-get install --no-install-recommends -y \
  software-properties-common \
  && rm -rf /var/lib/apt/lists/*

# Install gh
RUN apt-key adv --keyserver keyserver.ubuntu.com --recv-key C99B11DEB97541F0 && \
    apt-add-repository https://cli.github.com/packages && \
    apt-get update && apt-get install --no-install-recommends -y \
    gh \
    && rm -rf /var/lib/apt/lists/*

# Install standard version
RUN npm i -g standard-version

# Add release script
COPY release.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/release.sh

USER circleci

# Disable prompt
RUN gh config set prompt disabled

# Setup git
RUN git config --global user.email "alexrfalkowski@gmail.com" && \
    git config --global user.name "Alejandro Falkowski"
