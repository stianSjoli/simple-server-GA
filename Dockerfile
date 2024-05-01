FROM node:latest as buildFrontend
COPY package-lock.json package.json src/
COPY dist/index.html src/dist/index.html
COPY dist/style.css src/dist/style.css
COPY dist/robots.txt src/dist/robots.txt
COPY webpack.config.js src/webpack.config.js
COPY .babelrc src/
COPY src src/src

WORKDIR src/

RUN npm ci --legacy-peer-deps 
RUN npm run build
FROM nginx:1.26.0-alpine
COPY nginx.conf /etc/nginx/nginx.conf
RUN  apk update && apk add apparmor
USER 0
COPY --from=buildFrontend /src/dist/index.html /app/build/index.html
COPY --from=buildFrontend /src/dist/style.css /app/build/style.css
COPY --from=buildFrontend /src/dist/index.js /app/build/index.js
COPY --from=buildFrontend /src/dist/robots.txt /app/build/robots.txt
