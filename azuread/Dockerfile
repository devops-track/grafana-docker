FROM golang:1.16-alpine
RUN apk --no-cache add --update git make 
RUN git clone https://github.com/FXinnovation/azure-health-exporter.git

WORKDIR /go/azure-health-exporter
RUN make build \
    &&  ./azure-health-exporter
EXPOSE 9613
ENV AZURE_SUBSCRIPTION_ID=9ad70610-7927-49aa-b3af-2f4c1cc64699
ENV AZURE_TENANT_ID=a47bb6bd-ad61-40fc-bbcb-623fdc323c97
ENV AZURE_CLIENT_ID=8cb59d88-19aa-4827-bc91-35df15fc1460
ENV AZURE_CLIENT_SECRET=TJ~7Q~2sFuckmb.Rdi.gCJyvM9f4~28INbQbU