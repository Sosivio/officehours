FROM centos:7

RUN yum install -y  epel-release 
RUN yum install -y stress
COPY stress.sh /
RUN chmod +x /stress.sh
CMD /stress.sh
