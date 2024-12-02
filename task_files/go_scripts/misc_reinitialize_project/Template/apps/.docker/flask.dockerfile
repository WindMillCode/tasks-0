FROM [PROJECT_NAME]/base_img:latest

ENV PYENV_ROOT=/root/.pyenv
ENV PATH $PYENV_ROOT/bin:/root/.pyenv/plugins/xxenv-latest/bin:$PATH
RUN git clone https://github.com/pyenv/pyenv.git ~/.pyenv ;
RUN git clone https://github.com/momo-lab/xxenv-latest.git $(pyenv root)/plugins/xxenv-latest
RUN git clone https://github.com/pyenv/pyenv-update.git $(pyenv root)/plugins/pyenv-update;
RUN git clone https://github.com/pyenv/pyenv-virtualenv.git $(pyenv root)/plugins/pyenv-virtualenv
RUN echo 'eval "$(pyenv init -)"'  >> /root/.profile ; \
    echo 'eval "$(pyenv init --path)"' >> /root/.profile ; \
    echo 'eval "$(pyenv virtualenv-init -)"' >> /root/.profile \
    echo 'eval "$(pyenv init -)"'  >> /root/.bashrc ; \
    echo 'eval "$(pyenv init --path)"' >> /root/.bashrc ; \
    echo 'eval "$(pyenv virtualenv-init -)"' >> /root/.bashrc
RUN pyenv install 3.11.6  ;

COPY ./.secrets .
WORKDIR /FlaskApp
ARG     ENVIRON=docker-local

COPY ./backend/FlaskApp .

RUN eval "$(pyenv init --path)"; \
  eval "$(pyenv init -)"; \
  eval "$(pyenv virtualenv-init -)"; \
  pyenv shell 3.11.6; \
  pip install setuptools --upgrade;\
  pip install -r docker-requirements.txt --upgrade;

EXPOSE 5000

COPY .docker/flask_run_script.sh /bin/flask_run_script.sh
RUN chmod 777 /bin/flask_run_script.sh ;

ENTRYPOINT ["flask_run_script.sh"]
CMD ["python","app.py"]
