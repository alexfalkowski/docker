FROM circleci/node:15-buster

USER root

# Install standard version
RUN npm i -g standard-version

# Add release script
COPY release.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/release.sh

USER circleci

# Setup git
RUN git config --global user.email "alexrfalkowski@gmail.com" && \
    git config --global user.name "Alejandro Falkowski"
