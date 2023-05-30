#build golang image
FROM golang:1.20 as base

FROM base as dev

# Configure to reduce warnings and limitations as instruction from official VSCode Remote-Containers.
# See https://code.visualstudio.com/docs/remote/containers-advanced#_reducing-dockerfile-build-warnings.
ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update \
    && apt-get -y install --no-install-recommends apt-utils 2>&1

# Verify git, process tools, lsb-release (common in install instructions for CLIs) installed.
RUN apt-get -y install git iproute2 procps lsb-release

# Install Go tools.
RUN apt-get update \
    # Install gocode-gomod.
    # Install other tools.
    # && go get -v golang.org/x/tools/gopls \
    # && go get github.com/google/wire/cmd/wire \
    # Clean up.
    && apt-get autoremove -y \
    && apt-get clean -y \
    && rm -rf /var/lib/apt/lists/*

# Revert workaround at top layer.
ENV DEBIAN_FRONTEND=dialog

WORKDIR /server

EXPOSE 50051

ENTRYPOINT ["tail", "-f", "/dev/null"]
