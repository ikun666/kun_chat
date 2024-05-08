FROM busybox 
 
LABEL author="ikun666"
 
ENV TZ Asia/Shanghai
EXPOSE 8000
# WORKDIR /app
ADD userapi /
ADD etc/user.yaml /
CMD ["/userapi", "-f", "/user.yaml"]