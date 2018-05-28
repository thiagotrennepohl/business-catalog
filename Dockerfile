FROM scratch

ADD ./business-catalog /business-catalog

CMD ["/business-catalog"]