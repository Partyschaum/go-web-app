--
-- PostgreSQL database dump
--

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: member; Type: TABLE; Schema: public; Owner: lss_admin; Tablespace: 
--

CREATE TABLE member (
    id integer NOT NULL,
    email character varying,
    password character varying,
    first_name character varying
);


ALTER TABLE member OWNER TO lss_admin;

--
-- Name: member_id_seq; Type: SEQUENCE; Schema: public; Owner: lss_admin
--

CREATE SEQUENCE member_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE member_id_seq OWNER TO lss_admin;

--
-- Name: member_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: lss_admin
--

ALTER SEQUENCE member_id_seq OWNED BY member.id;


--
-- Name: session; Type: TABLE; Schema: public; Owner: lss_admin; Tablespace: 
--

CREATE TABLE session (
    id integer NOT NULL,
    member_id integer,
    session_id character varying
);


ALTER TABLE session OWNER TO lss_admin;

--
-- Name: session_id_seq; Type: SEQUENCE; Schema: public; Owner: lss_admin
--

CREATE SEQUENCE session_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE session_id_seq OWNER TO lss_admin;

--
-- Name: session_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: lss_admin
--

ALTER SEQUENCE session_id_seq OWNED BY session.id;


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: lss_admin
--

ALTER TABLE ONLY member ALTER COLUMN id SET DEFAULT nextval('member_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: lss_admin
--

ALTER TABLE ONLY session ALTER COLUMN id SET DEFAULT nextval('session_id_seq'::regclass);


--
-- Data for Name: member; Type: TABLE DATA; Schema: public; Owner: lss_admin
--

COPY member (id, email, password, first_name) FROM stdin;
1	hauke@jimdo.com	2bb80d537b1da3e38bd30361aa855686bde0eacd7162fef6a25fe97bf527a25b	Hauke
\.


--
-- Name: member_id_seq; Type: SEQUENCE SET; Schema: public; Owner: lss_admin
--

SELECT pg_catalog.setval('member_id_seq', 1, true);


--
-- Data for Name: session; Type: TABLE DATA; Schema: public; Owner: lss_admin
--

COPY session (id, member_id, session_id) FROM stdin;
4	1	92b61cb54173e538f3a1128c387785a0cfba8528afc7272865a5d6cf459eec8a
\.


--
-- Name: session_id_seq; Type: SEQUENCE SET; Schema: public; Owner: lss_admin
--

SELECT pg_catalog.setval('session_id_seq', 4, true);


--
-- Name: member_pkey; Type: CONSTRAINT; Schema: public; Owner: lss_admin; Tablespace: 
--

ALTER TABLE ONLY member
    ADD CONSTRAINT member_pkey PRIMARY KEY (id);


--
-- Name: session_pkey; Type: CONSTRAINT; Schema: public; Owner: lss_admin; Tablespace: 
--

ALTER TABLE ONLY session
    ADD CONSTRAINT session_pkey PRIMARY KEY (id);


--
-- Name: public; Type: ACL; Schema: -; Owner: hauke
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM hauke;
GRANT ALL ON SCHEMA public TO hauke;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--

