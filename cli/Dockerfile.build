FROM mhart/alpine-node:6
MAINTAINER Arachnys <techteam@arachnys.com>

RUN mkdir -p /athenapdf/build/artifacts/
WORKDIR /athenapdf/

COPY package.json /athenapdf/
RUN npm install

COPY package.json /athenapdf/build/artifacts/
RUN cp -r /athenapdf/node_modules/ /athenapdf/build/artifacts/

COPY src /athenapdf/build/artifacts/
RUN npm run build:linux

CMD ["/bin/sh"]