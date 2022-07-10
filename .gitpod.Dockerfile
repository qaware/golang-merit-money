FROM gitpod/workspace-full

RUN install-packages  \
        shellcheck && \
        golangci-lint && \
    pip3 install pre-commit