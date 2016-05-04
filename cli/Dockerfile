FROM debian:8.3
MAINTAINER Arachnys <techteam@arachnys.com>

RUN echo 'deb http://httpredir.debian.org/debian/ stable main contrib non-free' >> /etc/apt/sources.list

RUN apt-get -yq update && \
    apt-get -yq install \
        wget \
        xvfb \
        libgtk2.0-0 \
        libgconf-2-4 \
        libXtst6 \
        libnss3 \
        libnotify4 \
        culmus \
        fonts-beng \
        fonts-hosny-amiri \
        fonts-lklug-sinhala \
        fonts-lohit-guru \
        fonts-lohit-knda \
        fonts-samyak-gujr \
        fonts-samyak-mlym \
        fonts-samyak-taml \
        fonts-sarai \
        fonts-sil-abyssinica \
        fonts-sil-padauk \
        fonts-telu \
        fonts-thai-tlwg \
        ttf-liberation \
        ttf-unfonts-core \
        ttf-wqy-zenhei \
    && apt-get -yq autoremove \
    && apt-get -yq clean \
    && rm -rf /var/lib/apt/lists/* \
    && truncate -s 0 /var/log/*log

ADD build/athenapdf-linux-x64/ /athenapdf/
WORKDIR /athenapdf/

ENV PATH /athenapdf/:$PATH

ADD entrypoint.sh /athenapdf/entrypoint.sh

RUN mkdir -p /converted/
WORKDIR /converted/

CMD ["athenapdf"]

ENTRYPOINT ["/athenapdf/entrypoint.sh"]
