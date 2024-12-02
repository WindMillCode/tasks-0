
apt-get -y update && DEBIAN_FRONTEND=noninteractive;
apt-get -y install postgresql-16-cron

echo "shared_preload_libraries = 'pg_cron'" >> $PGDATA/postgresql.conf
echo "cron.database_name = 'postgres'" >> $PGDATA/postgresql.conf
