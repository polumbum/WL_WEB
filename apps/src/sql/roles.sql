CREATE ROLE guest WITH
NOSUPERUSER
NOCREATEDB
NOCREATEROLE
LOGIN
NOREPLICATION
PASSWORD 'guest'
CONNECTION LIMIT -1;

GRANT SELECT ON
public.competitions, 
public.t_camps,
public.results
TO guest;

GRANT INSERT ON
public.users, 
TO guest;


CREATE ROLE secretary WITH
NOSUPERUSER
NOCREATEDB
NOCREATEROLE
LOGIN
NOREPLICATION
PASSWORD 'secretary'
CONNECTION LIMIT -1;

GRANT SELECT ON
ALL TABLES IN SCHEMA public
TO secretary;

GRANT INSERT ON
public.comp_accesses,
public.antidopings,
public.sportsmen_coaches,
public.results
TO secretary;

GRANT UPDATE ON
public.sportsmen,
public.coaches,
public.comp_accesses,
public.antidopings,
public.sportsmen_coaches,
public.results,
public.competitions,
public.t_camps
TO secretary;

GRANT DELETE ON
public.sportsmen,
public.coaches,
public.comp_accesses,
public.antidopings,
public.sportsmen_coaches,
public.results,
public.competitions,
public.t_camps
TO secretary;


CREATE ROLE coach WITH
NOSUPERUSER
NOCREATEDB
NOCREATEROLE
LOGIN
NOREPLICATION
PASSWORD 'coach'
CONNECTION LIMIT -1;

GRANT SELECT ON
public.competitions, 
public.t_camps,
public.results,
public.comp_accesses,
public.antidopings,
public.sportsmen,
public.sportsmen_coaches,
public.coaches,
public.t_camp_applications,
public.comp_applications
TO coach;

GRANT INSERT ON
public.t_camp_applications,
public.comp_applications
TO coach;


CREATE ROLE sportsman WITH
NOSUPERUSER
NOCREATEDB
NOCREATEROLE
LOGIN
NOREPLICATION
PASSWORD 'sportsman'
CONNECTION LIMIT -1;

GRANT SELECT ON
public.competitions, 
public.t_camps,
public.results,
public.comp_accesses,
public.antidopings,
public.sportsmen,
public.sportsmen_coaches,
public.coaches,
public.t_camp_applications,
public.comp_applications
TO sportsman;

GRANT INSERT ON
public.t_camp_applications,
public.comp_applications
TO sportsman;


CREATE ROLE comp_org WITH
NOSUPERUSER
NOCREATEDB
NOCREATEROLE
LOGIN
NOREPLICATION
PASSWORD 'comp_org'
CONNECTION LIMIT -1;

GRANT SELECT ON
public.competitions, 
public.t_camps,
public.comp_applications
TO comp_org;

GRANT INSERT ON
public.competitions
TO comp_org;

GRANT UPDATE ON
public.competitions
TO comp_org;

GRANT DELETE ON
public.competitions
TO comp_org;


CREATE ROLE tcamp_org WITH
NOSUPERUSER
NOCREATEDB
NOCREATEROLE
LOGIN
NOREPLICATION
PASSWORD 'tcamp_org'
CONNECTION LIMIT -1;

GRANT SELECT ON
public.competitions, 
public.t_camps,
public.t_camp_applications
TO tcamp_org;

GRANT INSERT ON
public.t_camps
TO tcamp_org;

GRANT UPDATE ON
public.t_camps
TO tcamp_org;

GRANT DELETE ON
public.t_camps
TO tcamp_org;


SELECT * FROM pg_catalog.pg_roles;
SELECT * FROM information_schema.role_table_grants
where grantee = 'tcamp_org' 

