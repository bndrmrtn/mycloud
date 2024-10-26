FROM node:20.17.0

WORKDIR /app

COPY package*.json ./

RUN npm install -g npm@10.9.0
RUN npm install

COPY . .

EXPOSE 3000

ENV API_URL "/api"
RUN npm run build

CMD [ "npm", "run", "start" ]