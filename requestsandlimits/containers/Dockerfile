FROM ubuntu:20.10
RUN apt update -y
RUN apt install -y  stress-ng python3
COPY stress.sh /
COPY memeater.py /
RUN chmod +x /stress.sh
CMD /stress.sh
