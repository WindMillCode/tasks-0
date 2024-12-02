# TODO BEFORE RUNNING THE RESTORE SCRIPT CHECK THE TRUNCATE TABLE SQL IN THE BACKUP AND THE dump_db.sh always make sure its not including any schema that are not yours
#!/bin/bash

DB_NAME="[PROJECT_NAME]-postgres-db-0"
DB_USER="postgres"

# Ensure that all necessary positional parameters are provided
if [ "$#" -lt 3 ]; then
  echo "Usage: $0 restoreFileType schemaOnly backupFileName"
  exit 1
fi

# Assign positional parameters to variables
RESTORE_FILE_TYPE=$1
SCHEMA_ONLY=$2
MY_BACKUP_FILE=$3



RESTORE_FILE_TYPE=$(if [ "$RESTORE_FILE_TYPE" == "SQL" ]; then echo "t"; else echo "c"; fi)
SCHEMA_ONLY=$(if [ "$SCHEMA_ONLY" == "TRUE" ]; then echo "--schema-only"; else echo ""; fi)

if [ "$RESTORE_FILE_TYPE" == "t" ]; then
  RESTORE_CMDS=(
    "psql -U $DB_USER -d $DB_NAME   -v AUTOCOMMIT=on -f $MY_BACKUP_FILE"
  )
else
  RESTORE_CMDS=(
    "pg_restore -C -U $DB_USER -d $DB_NAME $SCHEMA_ONLY -F $RESTORE_FILE_TYPE $MY_BACKUP_FILE"
  )
fi

for r_cmd in "${RESTORE_CMDS[@]}"; do
  echo "$r_cmd"
  eval $r_cmd
done




