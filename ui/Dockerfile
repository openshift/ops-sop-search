FROM node:12

WORKDIR /usr/source/ui

COPY . .

EXPOSE 9000

RUN yarn global add serve

ENTRYPOINT serve -s build