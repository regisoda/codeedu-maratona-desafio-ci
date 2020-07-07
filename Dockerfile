FROM scratch

COPY gopath/bin/src /src

ENTRYPOINT ["/src"]