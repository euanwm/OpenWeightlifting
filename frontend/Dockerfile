# base image
FROM node:18.7.0
ARG API
ENV API $API

WORKDIR /usr/src/app

# install and cache app dependencies
COPY package.json package-lock.json .
RUN npm install

# copy the rest of the files
COPY . .

# build
RUN --mount=type=cache,target=/usr/src/app/.next/cache npm run build

# start app
CMD ["npm", "start"]
