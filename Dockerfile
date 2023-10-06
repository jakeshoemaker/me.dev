FROM golang:latest
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build server/main.go

FROM node:latest
WORKDIR /app
COPY main.css . 
COPY package.json .
COPY tailwind.config.js .
RUN npm install
RUN npx tailwindcss build main.css -o server/controllers/static/styles.css

FROM golang:latest
WORKDIR /app
COPY --from=0 /app/main .
COPY --from=1 /app/main.css .
EXPOSE 8080
CMD ["./main"]
