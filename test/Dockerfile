FROM node:16-slim

WORKDIR /athenapdf-tests

COPY . /athenapdf-tests

RUN apt update && \
    apt install imagemagick graphicsmagick -y && \
    sed -i '/disable ghostscript format types/,+6d' /etc/ImageMagick-6/policy.xml && \
    npm install

