FROM cimg/openjdk:17.0.3

USER root

# Install deps
RUN curl -s https://packagecloud.io/install/repositories/github/git-lfs/script.deb.sh | bash && \
    apt-get update && apt-get install --no-install-recommends -y \
    git-lfs \
    python3 \
    python3-pip \
    python3-setuptools \
    graphviz \
    && rm -rf /var/lib/apt/lists/*

# Install lfs
RUN git lfs install

WORKDIR /usr/local/bin

# Install structurizr
RUN curl -o structurizr-cli.zip -L https://github.com/structurizr/cli/releases/download/v1.18.0/structurizr-cli-1.18.0.zip && \
	unzip -o structurizr-cli.zip && \
	rm -f structurizr-cli.zip

# Install plantuml
RUN curl -o plantuml.jar -L https://netix.dl.sourceforge.net/project/plantuml/plantuml.jar

USER circleci
WORKDIR /home/circleci/

# Install diagrams
RUN pip3 install --no-cache-dir -U diagrams==0.21.1 --user
