FROM alpine

COPY ./Globalizer /usr/loca/bin/Globalizer

ENTRYPOINT [ "/usr/loca/bin/Globalizer" ]