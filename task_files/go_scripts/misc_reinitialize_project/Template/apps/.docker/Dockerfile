FROM windmillcode/docker-img-0:0.0.18

WORKDIR /root

ENV NODE_OPTIONS="--max-old-space-size=8192"

RUN mkcert example.com "*.example.com" example.test localhost 127.0.0.1 ::1;\
    mv example.com+5.pem /root/cert.pem; mv example.com+5-key.pem /root/key.pem;

ENV WML_CERT0=/root/cert.pem
ENV WML_CERT_KEY0=/root/key.pem

COPY . ./[PROJECT_NAME_CAPITAL]

RUN eval "$(pyenv init --path)"; \
  eval "$(pyenv init -)"; \
  eval "$(pyenv virtualenv-init -)"; \
  pyenv global $(pyenv latest 3); \
  pyenv shell $(pyenv latest 3);

COPY ./apps/.docker/run_script.sh /bin/run_script
RUN chmod 777 /bin/run_script ;\
  dos2unix /bin/run_script




