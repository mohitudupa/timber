FROM alpine:3.14.2

# Copy linux build into /bin
COPY builds/timber-linux /bin/timber
RUN chmod 755 /bin/timber

# Copy default config
COPY timberconf.json /application/timberconf.json

# Setup environment
WORKDIR /application
EXPOSE 36036
ENTRYPOINT [ "timber" ]
