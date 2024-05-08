FROM busybox 
 
LABEL author="ikun666"
 
ENV TZ Asia/Shanghai
EXPOSE 8100
# WORKDIR /app
ADD relationapi /
ADD etc/relation.yaml /
CMD ["/relationapi", "-f", "/relation.yaml"]