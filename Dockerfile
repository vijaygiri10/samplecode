FROM jetsend_opens_base:latest as builder
#Declare Variables for path
ENV SERVICE_NAME services/jetsend_opens
ENV PKG_PATH /etc/jetsend

#Add service code to container
ADD . $PKG_PATH/$SERVICE_NAME

WORKDIR $PKG_PATH/$SERVICE_NAME

# Build Service as Executable:
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -ldflags="-w -s" main.go
############################
#STEP 2 build a small image
############################
FROM alpine

ENV SERVICE_NAME /services/jetsend_opens
ENV PKG_PATH /etc/jetsend
ENV logName jetsend_opens
ENV JETSEND_ENV uat

WORKDIR /

RUN apk update

# Import CA certs and TZ-Data from builder.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /etc/timezone /etc/timezone
WORKDIR $SERVICE_NAME

#Copying Go App Binary
COPY --from=builder $PKG_PATH/$SERVICE_NAME/main $SERVICE_NAME/
# Add Configuration file
COPY --from=builder $PKG_PATH/$SERVICE_NAME/config $SERVICE_NAME/config

RUN mkdir log

#Starting Go Application
CMD $SERVICE_NAME/main ./config ./log
