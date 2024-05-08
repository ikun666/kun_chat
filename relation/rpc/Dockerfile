FROM busybox 
 
LABEL author="ikun666"
 
ENV TZ Asia/Shanghai
EXPOSE 9100
# WORKDIR /app
ADD relationrpc /
ADD etc/relation.yaml /
CMD ["/relationrpc", "-f", "/relation.yaml"]