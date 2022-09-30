-- DROP SCHEMA public;

CREATE SCHEMA public AUTHORIZATION jamesfanuel;
-- public.friend_request definition

-- Drop table

-- DROP TABLE public.friend_request;

CREATE TABLE public.friend_request (
	requestor varchar NULL,
	"to" varchar NULL,
	status varchar NULL
);


-- public.user_friend_list definition

-- Drop table

-- DROP TABLE public.user_friend_list;

CREATE TABLE public.user_friend_list (
	username varchar NULL,
	friend varchar NULL,
	is_blocked int4 NULL DEFAULT 0
);