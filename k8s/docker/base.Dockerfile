FROM scratch
ADD ./scylla-cloud-linux /scylla-cloud
ENTRYPOINT ["/scylla-cloud"]
