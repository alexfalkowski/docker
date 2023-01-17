FROM cimg/go:1.19.5

USER root

# Install uplift.
RUN curl https://raw.githubusercontent.com/gembaadvantage/uplift/main/scripts/install | bash -s -- -v v2.20.0 --no-sudo

# Add release script.
COPY release /usr/local/bin/
RUN chmod +x /usr/local/bin/release

# Add uplift config.
RUN mkdir /etc/uplift
COPY .uplift.yml /etc/uplift

USER circleci

# Setup git.
RUN git config --global user.email "alexrfalkowski@gmail.com" && \
    git config --global user.name "Alejandro Falkowski"
