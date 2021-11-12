FROM centos:7

COPY ./main  /

ENTRYPOINT ["main"]
