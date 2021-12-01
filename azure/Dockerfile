FROM golang:1.16-alpine
RUN apk --no-cache add --update git make 
RUN git clone https://github.com/RobustPerception/azure_metrics_exporter.git 

RUN rm -rf /go/azure_metrics_exporter/azure-example.yml
ADD az.yml /go/azure_metrics_exporter/ 
RUN mv /go/azure_metrics_exporter/az.yml /go/azure_metrics_exporter/azure.yml
WORKDIR /go/azure_metrics_exporter
RUN make build \
    &&  ./azure_metrics_exporter
EXPOSE 9276