FROM cimg/openjdk:19.0.2

USER root

# Install deps
RUN apt-get update && apt-get install --no-install-recommends -y \
    python3 \
    python3-pip \
    python3-setuptools \
    graphviz \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /usr/local/bin

# Install structurizr
RUN curl -o structurizr-cli.zip -L https://github.com/structurizr/cli/releases/download/v1.20.1/structurizr-cli-1.20.1.zip && \
    unzip -o structurizr-cli.zip && \
    rm -f structurizr-cli.zip

# Install plantuml
RUN curl -o plantuml.jar -L https://netix.dl.sourceforge.net/project/plantuml/plantuml.jar

USER circleci
WORKDIR /home/circleci/

# Install diagrams
RUN pip3 install --no-cache-dir -U diagrams==0.22.0 --user
