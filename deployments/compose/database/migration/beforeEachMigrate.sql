DO $$
 BEGIN
    IF EXISTS (SELECT rolname FROM pg_roles where rolname like 'hamburgueria') THEN
	 SET ROLE "hamburgueria";
	end if;
  end;
$$;

SET STATEMENT_TIMEOUT TO '300s';
