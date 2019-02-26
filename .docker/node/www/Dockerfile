# FROM node:10.13.0-alpine
# ENV NODE_ENV=development
# ENV HOST 0.0.0.0 
# # for Nginx reverse proxy, NUXT must host on 0.0.0.0

# RUN mkdir -p /var/www/html

# # COPY . /var/www/html

# WORKDIR /var/www

# COPY ./package.json /var/www/package.json
# RUN npm install
# RUN npm run build
# RUN npm install pm2 -g
# COPY ./process.yml /var/www/process.yml
# CMD pm2-dev start process.yml


# for build
FROM node:10.13.0-alpine as builder

ENV NODE_ENV=production

RUN mkdir -p /app
COPY . /app
WORKDIR /app

RUN npm run build

# =============================

# for start
FROM node:10.13.0-alpine
LABEL maintainer="junxiang"

ENV HOST=0.0.0.0

RUN mkdir -p /app
WORKDIR /app

ADD package.json ./
ADD nuxt.config.js ./

COPY --from=builder ./app/server ./server/
COPY --from=builder ./app/node_modules ./node_modules/
COPY --from=builder ./app/.nuxt ./.nuxt/
COPY --from=builder ./app/static ./static/

EXPOSE 3000

CMD ["npm", "run", "start"]