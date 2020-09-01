-- This is probably not the best way to handle unique IDs, but the data provided used regular integers 
CREATE SEQUENCE idSeq START WITH 500;
ALTER TABLE customer ALTER column id SET DEFAULT netxval('idSet');

CREATE SEQUENCE midSeq START WITH 250;
ALTER TABLE mechanic ALTER COLUMN id SET DEFAULT nextval('midSeq');

CREATE SEQUENCE oidSeq START WITH 5000;
ALTER TABLE owns ALTER COLUMN ownership_id SET DEFAULT nextval('oidSeq');

CREATE SEQUENCE ridSeq START WITH 30001;
ALTER TABLE service_request ALTER COLUMN rid SET DEFAULT nextval('ridSeq');

CREATE SEQUENCE widSeq START WITH 30001;
ALTER TABLE closed_request ALTER COLUMN wid SET DEFAULT nextval('widSeq');