FROM busybox 
 
LABEL author="ikun666"
 
ENV TZ Asia/Shanghai
EXPOSE 8200
# WORKDIR /app
ADD chatapi /
ADD etc/chat.yaml /
CMD ["/chatapi", "-f", "/chat.yaml"]