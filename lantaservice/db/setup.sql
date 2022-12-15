CREATE TABLE IF NOT EXISTS public.news
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    title character varying COLLATE pg_catalog."default",
    text character varying COLLATE pg_catalog."default",
    date character varying COLLATE pg_catalog."default",
    CONSTRAINT news_pkey PRIMARY KEY (id)
)
