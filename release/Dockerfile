FROM circleci/node:14-buster

# Install standard version
USER root
RUN npm i -g standard-version
COPY release.sh /usr/local/bin/

# Setup git
USER circleci
RUN git config --global user.email "alexrfalkowski@gmail.com" && \
    git config --global user.name "Alejandro Falkowski"
