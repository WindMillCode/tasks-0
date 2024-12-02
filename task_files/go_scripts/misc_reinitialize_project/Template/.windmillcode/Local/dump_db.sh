#!/bin/bash

DB_NAME="[PROJECT_NAME]-postgres-db-0"
DB_USER="postgres"


# Ensure that all necessary positional parameters are provided
if [ "$#" -lt 3 ]; then
  echo "Usage: $0 dumpFileType schemaOnly backupFileName"
  exit 1
fi

# Assign positional parameters to variables
DUMP_FILE_TYPE=$1
SCHEMA_ONLY=$2
MY_BACKUP_FILE=$3




DUMP_FILE_TYPE=$(if [ "$DUMP_FILE_TYPE" == "SQL" ]; then echo "p"; else echo "c"; fi)
SCHEMA_ONLY=$(if [ "$SCHEMA_ONLY" == "TRUE" ]; then echo "--schema-only"; else echo " --column-inserts"; fi)



# List all schemas with a specific comment
SCHEMAS=$(psql -U $DB_USER -d $DB_NAME -t -c "
  SELECT schema_name
  FROM information_schema.schemata
  WHERE catalog_name = '[PROJECT_NAME]-postgres-db-0'
  AND schema_name NOT IN ('information_schema', 'pg_catalog')
  AND schema_name NOT LIKE 'pg_toast%'
  AND schema_name NOT LIKE 'pg_temp%';
  ")

# Construct the pg_dump command
DUMP_CMD="pg_dump -U $DB_USER -d $DB_NAME"

for SCHEMA in $SCHEMAS; do
  DUMP_CMD+=" -n $SCHEMA"
done

DUMP_CMD+="$SCHEMA_ONLY -F $DUMP_FILE_TYPE -f $MY_BACKUP_FILE"
echo $DUMP_CMD
# Execute the dump command
eval $DUMP_CMD

if [ "$DUMP_FILE_TYPE" == "p" ]; then
  # disable checks and truncate tables to properly restore the db

  echo "
SET session_replication_role = replica;
SET constraint_checks = false;
DO $$
DECLARE
  r RECORD;
BEGIN
  FOR r IN (SELECT schemaname, tablename FROM pg_tables
            WHERE schemaname NOT IN ('information_schema', 'pg_catalog')
              AND schemaname NOT LIKE 'pg_toast%'
              AND schemaname NOT LIKE 'pg_temp%')
  LOOP
    EXECUTE 'TRUNCATE TABLE ' || quote_ident(r.schemaname) || '.' || quote_ident(r.tablename) || ' CASCADE';
  END LOOP;
END $$;
SET constraint_checks = true;
" > ${MY_BACKUP_FILE}_tmp
  cat $MY_BACKUP_FILE >> ${MY_BACKUP_FILE}_tmp
  echo 'SET session_replication_role = DEFAULT;' >> ${MY_BACKUP_FILE}_tmp

  mv ${MY_BACKUP_FILE}_tmp $MY_BACKUP_FILE
fi
