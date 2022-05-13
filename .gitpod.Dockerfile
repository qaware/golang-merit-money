FROM gitpod/workspace-full

USER root

RUN install-packages  \
        shellcheck && \
        golangci-lint && \
    pip3 install pre-commit && \
    chown -R gitpod:gitpod /home/gitpod/.cache/golangci-lint