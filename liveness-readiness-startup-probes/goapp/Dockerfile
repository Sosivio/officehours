FROM centos AS s0
RUN useradd gouser
RUN mkdir /assets
COPY assets/ /assets/
RUN chmod 775 -R /assets 

USER gouser
ADD goapp /
ADD start.sh /
CMD ["/start.sh"]
