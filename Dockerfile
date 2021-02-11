FROM golang:1.13.6 as review-central
ADD . review-central
RUN cd review-central; rm go.sum; go clean -modcache; go build -o review-central-build . ; mv review-central-build /

FROM node as ui
ADD ui ui
RUN cd ui; npm install --only=production; npm run build; mv build /

FROM ubuntu
RUN apt-get update; apt-get install -y ca-certificates; update-ca-certificates
COPY --from=review-central /review-central-build /app/review-central-build
COPY --from=ui /build /app/ui/build

WORKDIR /app
CMD ./review-central-build
