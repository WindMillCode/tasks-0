FROM [PROJECT_NAME]/base_img:latest

ENV NVM_DIR="/root/.nvm"
RUN curl -o- https://raw.githubusercontent.com/creationix/nvm/v0.39.5/install.sh | bash && \
    export NVM_DIR="$HOME/.nvm" && [ -s "$NVM_DIR/nvm.sh" ] && . "$NVM_DIR/nvm.sh" && \
    nvm install lts/hydrogen && \
	  nvm use lts/hydrogen && \
    npm i -g npm && \
    npm install -g yarn

RUN chmod -R 777 /root/.nvm;

USER root
WORKDIR /

RUN echo "ubuntu ALL=(ALL:ALL) ALL" >> /etc/sudoers.d/ubuntu
RUN    [ -s "$NVM_DIR/nvm.sh" ]  ; \
  . "$NVM_DIR/nvm.sh"  ;



WORKDIR /AngularApp
ARG     ENVIRON=docker-local

COPY ./frontend/AngularApp .
# RUN   rm -r  node_modules; \
#   rm -r coverage; \
#   rm -r .angular; \
#   rm -r dist; \
#   rm  package-lock.json;


RUN    [ -s "$NVM_DIR/nvm.sh" ]  ; \
  . "$NVM_DIR/nvm.sh"  ;\
  # cp package.json ${NVM_DIR}/versions/node/$(node -v)/lib;\
  # cd ${NVM_DIR}/versions/node/$(node -v)/lib;\
  npm install   --force --verbose;



COPY .docker/angular_run_script.sh /bin/angular_run_script.sh
RUN chmod 777 /bin/angular_run_script.sh ;

EXPOSE [Angular_Run_0]


ENTRYPOINT ["angular_run_script.sh"]
# CMD ["npx", "ng", "serve", "-c", "docker-dev", "--ssl", "true", "--ssl-key", "$WML_CERT_KEY0", "--ssl-cert", "$WML_CERT0"]
# CMD ["npx ng serve -c docker-dev --ssl true --ssl-key $WML_CERT_KEY0 --ssl-cert $WML_CERT0"]
