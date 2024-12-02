DECLARE
  my_schemas text[];
BEGIN
  SELECT array_agg(schema_name)
  INTO my_schemas
  FROM information_schema.schemata
  WHERE catalog_name = '[PROJECT_NAME]-postgres-db-0'
  AND schema_name NOT IN ('information_schema', 'pg_catalog')
  AND schema_name NOT LIKE 'pg_toast%'
  AND schema_name NOT LIKE 'pg_temp%';

  FOREACH my_schema IN ARRAY my_schemas
  LOOP
    EXECUTE format('COMMENT ON SCHEMA %I IS ''created_by_[PROJECT_NAME]''', my_schema);
  END LOOP;
END
