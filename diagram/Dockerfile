FROM circleci/openjdk:16-buster-node-browsers-legacy

USER root
WORKDIR /usr/local/bin

# Install structurizr
RUN curl -o structurizr-cli.zip -L https://github.com/structurizr/cli/releases/download/v1.18.0/structurizr-cli-1.18.0.zip && \
	unzip -o structurizr-cli.zip && \
	rm -f structurizr-cli.zip

# Install plantuml
RUN curl -o plantuml.jar -L https://netix.dl.sourceforge.net/project/plantuml/plantuml.jar

# Install python3
RUN apt-get update && apt-get install --no-install-recommends -y \
  python3 \
  python3-pip \
  graphviz \
  && rm -rf /var/lib/apt/lists/*

USER circleci
WORKDIR /home/circleci/

# Install diagrams
RUN pip3 install --no-cache-dir -U diagrams==0.20.0 --user
