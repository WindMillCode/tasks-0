#!/bin/bash

function initialize_postgres_and_ssh() {
    echo "Waiting for PostgreSQL to start..."
    until pg_isready -h localhost -p 5432; do
        sleep 1
    done
    echo "PostgreSQL started."

    cd ~/[PROJECT_NAME_CAPITAL]/apps/database/postgres
    psql -U postgres -d template1 -a -f init_mysql_db_conn.sql

    eval "$(ssh-agent -s)"
    cp ~/[PROJECT_NAME_CAPITAL]/apps/.secrets/[VCS_PRIVATE_KEY] ~/.ssh/[VCS_PRIVATE_KEY]
    chmod 600 ~/.ssh/[VCS_PRIVATE_KEY]
    dos2unix ~/.ssh/[VCS_PRIVATE_KEY]
    ssh-add ~/.ssh/[VCS_PRIVATE_KEY]
    echo '   IdentityFile ~/.ssh/[VCS_PRIVATE_KEY]' >> /etc/ssh/ssh_config
    export WINDMILLCODE_DEV_CONTAINER_INIT=1
}

if [ -z "$WINDMILLCODE_DEV_CONTAINER_INIT" ] || [ "$WINDMILLCODE_DEV_CONTAINER_INIT" != "1" ]; then
    echo "Environment variable either does not exist or is not equal to 1. Executing your code..."
    export NVM_DIR=/root/.nvm
    export JAVA_HOME=/root/.jabba/jdk/openjdk@1.17.0
    export PYENV_ROOT=/root/.pyenv
    export PATH=/usr/lib/postgresql/16/bin:$NVM_DIR:$JAVA_HOME:$PYENV_ROOT:$PATH
    export PGPASSWORD=mysecretpassword
    export POSTGRES_PASSWORD=$PGPASSWORD
    export PGDATA=/var/lib/postgresql/data
    export POSTGRES_INITDB_WALDIR=$PGDATA

    if [ ! -f ~/.bashrc ]; then
        touch ~/.bashrc
    fi
    echo ' export PS1='$(basename "$(pwd)")> '' >> ~/.bashrc ;
    echo ' export PS1='$(basename "$(pwd)")> '' >> ~/.profile ;


    [ -s "$NVM_DIR/nvm.sh" ]
    . $NVM_DIR/nvm.sh


    eval "$(pyenv init -)"
    eval "$(pyenv virtualenv-init -)"

    pyenv global $(pyenv latest 3);
    pyenv shell $(pyenv latest 3);

    initialize_postgres_and_ssh &
    source /usr/local/bin/postgres_setup.sh;
    _main postgres;
    pg_ctl -D /var/lib/postgresql/data -l logfile start &


else
    echo "Environment variable is equal to 1. Skipping execution."
fi
