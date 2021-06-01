FROM alpine
ADD etcd /usr/local/bin/
ADD etcdctl /usr/local/bin/

RUN mkdir -p /var/etcd/
RUN mkdir -p /var/lib/etcd/

RUN echo 'hosts: files mdns4_minimal [NOTFOUND=return] dns mdns4' >> /etc/nsswitch.conf

EXPOSE 2379 2380

CMD [ "/usr/local/bin/etcd" ]
