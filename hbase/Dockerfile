FROM openjdk:18-jdk-buster

USER root
WORKDIR /tmp

# Install hbase.
RUN curl -L https://downloads.apache.org/hbase/2.4.9/hbase-2.4.9-bin.tar.gz > hbase-2.4.9-bin.tar.gz && \
  mkdir /usr/local/hbase-2.4.9 && \
  tar -xzvf hbase-2.4.9-bin.tar.gz -C /usr/local/hbase-2.4.9 --strip-components=1 && \
  rm -rf hbase-2.4.9-bin.tar.gz

# Install phoenix.
RUN curl -L https://downloads.apache.org/phoenix/phoenix-5.1.2/phoenix-hbase-2.4-5.1.2-bin.tar.gz > phoenix-hbase-2.4-5.1.2-bin.tar.gz && \
  mkdir apache-phoenix && \
  tar -xzvf phoenix-hbase-2.4-5.1.2-bin.tar.gz -C apache-phoenix --strip-components=1 && \
  cp apache-phoenix/phoenix-server-hbase-2.4-5.1.2.jar /usr/local/hbase-2.4.9/lib && \
  rm -rf apache-phoenix phoenix-hbase-2.4-5.1.2-bin.tar.gz

ENV PATH "$PATH:/usr/local/hbase-2.4.9/bin"

USER nobody
CMD ["hbase", "master", "start"]
