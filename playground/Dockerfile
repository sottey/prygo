FROM golang:alpine AS base

RUN apk add --no-cache binutils git subversion mercurial nodejs npm icu-data-full
RUN strip /usr/local/go/bin/go
RUN rm /usr/local/go/bin/gofmt
RUN rm -r /usr/local/go/src/cmd
#RUN rm -r /usr/local/go/src/vendor
RUN rm -r /usr/local/go/api
RUN cp /usr/local/go/lib/wasm/wasm_exec.js /tmp/wasm_exec.js
RUN rm -r /usr/local/go/lib
RUN rm -r /usr/local/go/misc
RUN rm -r /usr/local/go/test
RUN rm -r /usr/share
RUN find /usr/local/go/src -name "testdata" -exec rm -r {} +
RUN if [ -d /usr/lib/python2.7 ]; then find /usr/lib/python2.7 -name "*.pyo" -exec rm -r {} +; fi
RUN if [ -d /usr/lib/python2.7 ]; then find /usr/lib/python2.7 -name "*.pyc" -exec rm -r {} +; fi
RUN if [ -d /usr/local/go/pkg/linux_amd64 ]; then rm -r /usr/local/go/pkg/linux_amd64; fi
RUN strip /usr/local/go/pkg/tool/linux_amd64/*

WORKDIR /go/src/app
COPY . .
WORKDIR /go/src/app/playground

RUN npm install --omit=dev --ignore-scripts
RUN mv /tmp/wasm_exec.js ./wasm_exec.js

RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o playground ./server
RUN rm -rf /go/src/github.com /go/src/golang.org /go/bin
RUN rm -r /root

EXPOSE 8080
CMD ["./playground"]

#WORKDIR /go/src/app
#COPY . .
#
#FROM scratch
#ENV LD_LIBRARY_PATH /lib/:/usr/lib/
#
#COPY --from=base /go/src/app/ /
#
#COPY --from=base /usr/local/go/bin/go /bin/
#COPY --from=base /usr/local/go/src /usr/local/go/
#
#COPY --from=base /lib/*.so* /lib/
#COPY --from=base /usr/lib/*.so* /usr/lib/
#
#COPY --from=base /usr/bin/git /bin/
#COPY --from=base /usr/bin/svn /bin/
#COPY --from=base /usr/bin/hg /bin/
#
#COPY --from=base /tmp/ /tmp/

#EXPOSE 8080
#CMD ["/playground"]
