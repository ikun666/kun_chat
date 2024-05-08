FROM busybox 
 
LABEL author="ikun666"
 
ENV TZ Asia/Shanghai
EXPOSE 9000
# WORKDIR /app
ADD userrpc /
ADD etc/user.yaml /
CMD ["/userrpc", "-f", "/user.yaml"]