FROM openjdk:8-jdk-buster

USER root
WORKDIR /tmp

# Install hbase.
RUN curl -L https://downloads.apache.org/hbase/2.4.4/hbase-2.4.4-bin.tar.gz > hbase-2.4.4-bin.tar.gz && \
  mkdir /usr/local/hbase-2.4.4 && \
  tar -xzvf hbase-2.4.4-bin.tar.gz -C /usr/local/hbase-2.4.4 --strip-components=1 && \
  rm -rf hbase-2.4.4-bin.tar.gz

# Install phoenix.
RUN curl -L https://downloads.apache.org/phoenix/apache-phoenix-5.0.0-HBase-2.0/bin/apache-phoenix-5.0.0-HBase-2.0-bin.tar.gz > apache-phoenix-5.0.0-HBase-2.0-bin.tar.gz && \
  mkdir apache-phoenix && \
  tar -xzvf apache-phoenix-5.0.0-HBase-2.0-bin.tar.gz -C apache-phoenix --strip-components=1 && \
  cp apache-phoenix/phoenix-5.0.0-HBase-2.0-server.jar /usr/local/hbase-2.4.4/lib && \
  rm -rf apache-phoenix apache-phoenix-5.0.0-HBase-2.0-bin.tar.gz

ENV PATH "$PATH:/usr/local/hbase-2.4.4/bin"

USER nobody
CMD ["hbase", "master", "start"]
