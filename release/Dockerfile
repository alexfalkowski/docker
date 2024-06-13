FROM alexfalkowski/root:1.0

USER root

# Install uplift.
RUN echo 'deb [trusted=yes] https://fury.upliftci.dev/apt/ /' | tee /etc/apt/sources.list.d/uplift.list
RUN apt-get update && apt-get install --no-install-recommends -y \
    uplift=2.23.0 \
    && apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

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
