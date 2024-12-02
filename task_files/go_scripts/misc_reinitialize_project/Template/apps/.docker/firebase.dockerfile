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

# java
RUN curl -sL https://github.com/shyiko/jabba/raw/master/install.sh | \
    JABBA_COMMAND="install openjdk@1.17.0" bash

ENV JAVA_HOME=/root/.jabba/jdk/openjdk@1.17.0
ENV PATH $JAVA_HOME/bin:$PATH

# maven
RUN curl -s https://bitbucket.org/mjensen/mvnvm/raw/master/mvn > /bin/mvn && \
    chmod 0755 /bin/mvn

WORKDIR /FirebaseApp

COPY ./cloud/FirebaseApp .

RUN    [ -s "$NVM_DIR/nvm.sh" ]  ; \
  . "$NVM_DIR/nvm.sh"  ;\
  # cp package.json ${NVM_DIR}/versions/node/$(node -v)/lib/node_modules;\
  # cd ${NVM_DIR}/versions/node/$(node -v)/lib/node_modules;\
  npm install  --verbose --force;

COPY .docker/firebase_run_script.sh /bin/firebase_run_script.sh
RUN chmod 777 /bin/firebase_run_script.sh ;

ENTRYPOINT ["firebase_run_script.sh"]
CMD ["npx","firebase","emulators:start","--import=devData","--export-on-exit"]
