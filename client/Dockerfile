FROM node:11.15.0-alpine

WORKDIR '/app'

COPY *.json ./

RUN npm install 

COPY . .

CMD ["npm","start"]