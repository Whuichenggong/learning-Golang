CREATE DATABASE userinfo
(
    uid serial NOT NULL
    username character varying(100) NOT NULL,
    Create date.
    CONSTARAINT userinfo_pkey PRIMARY KEY (uid),
)
WITH (OIDS=FALSE);

CREATE TABLE userdetail
(
	uid integer,
	intro character varying(100),
	profile character varying(100)
)
WITH(OIDS=FALSE);