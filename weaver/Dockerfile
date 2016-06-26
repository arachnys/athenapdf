FROM arachnysdocker/athenapdf
MAINTAINER Arachnys <techteam@arachnys.com>

ENV GIN_MODE release

RUN \
  wget https://github.com/Yelp/dumb-init/releases/download/v1.0.0/dumb-init_1.0.0_amd64.deb \
  && dpkg -i dumb-init_*.deb \
  && rm dumb-init_*.deb \
  && mkdir -p /athenapdf-service/tmp/

COPY build/weaver /athenapdf-service/
WORKDIR /athenapdf-service/

ENV PATH /athenapdf-service/:$PATH

COPY conf/ /athenapdf-service/conf/

EXPOSE 8080

CMD ["dumb-init", "weaver"]

ENTRYPOINT ["/athenapdf-service/conf/entrypoint.sh"]