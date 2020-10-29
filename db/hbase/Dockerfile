FROM circleci/openjdk:8

USER root

# Install hbase.
RUN cd /tmp && \
  curl -L https://downloads.apache.org/hbase/2.3.2/hbase-2.3.2-bin.tar.gz > hbase-2.3.2-bin.tar.gz && \
  mkdir /usr/local/hbase-2.3.2 && \
  tar -xzvf hbase-2.3.2-bin.tar.gz -C /usr/local/hbase-2.3.2 --strip-components=1 && \
  rm -rf hbase-2.3.2-bin.tar.gz

ENV PATH "$PATH:/usr/local/hbase-2.3.2/bin"
