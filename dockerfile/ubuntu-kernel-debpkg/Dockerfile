FROM ubuntu:16.04

RUN sed -i 's/# deb-src/deb-src/g' /etc/apt/sources.list && \
    apt update

RUN apt-get install -y dpkg-dev debhelper gawk
RUN apt-get build-dep -y --no-install-recommends linux-image-4.13.0-36-generic

RUN apt-get source -y linux-image-4.13.0-36-generic

RUN cd linux-hwe-4.13.0/ && \
    touch build.sh && \
    chmod a+x build.sh && \
    printf '#! /bin/bash\n\
fakeroot debian/rules clean\n\
DEB_BUILD_OPTIONS=parallel=2 AUTOBUILD=1 fakeroot debian/rules binary-generic skipdbg=false\n'\
>> build.sh
