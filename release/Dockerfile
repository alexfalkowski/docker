FROM cimg/node:19.3.0

USER root

# Install deps
RUN curl -fsSL https://cli.github.com/packages/githubcli-archive-keyring.gpg | dd of=/usr/share/keyrings/githubcli-archive-keyring.gpg \
    && chmod go+r /usr/share/keyrings/githubcli-archive-keyring.gpg \
    && echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/githubcli-archive-keyring.gpg] https://cli.github.com/packages stable main" | tee /etc/apt/sources.list.d/github-cli.list > /dev/null \
    && apt-get update && apt-get install --no-install-recommends -y \
    software-properties-common \
    gh \
    && rm -rf /var/lib/apt/lists/*

# Setup deps
RUN npm i -g standard-version

# Add release script
COPY release /usr/local/bin/
RUN chmod +x /usr/local/bin/release

USER circleci

# Disable prompt
RUN gh config set prompt disabled

# Setup git
RUN git config --global user.email "alexrfalkowski@gmail.com" && \
    git config --global user.name "Alejandro Falkowski"
