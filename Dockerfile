FROM scratch
WORKDIR /app
ADD . /app
EXPOSE 8080
CMD ["/app/reminder"]
