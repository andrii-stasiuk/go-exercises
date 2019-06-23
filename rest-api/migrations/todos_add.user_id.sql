ALTER TABLE public.todos
ADD COLUMN user_id bigserial;

ALTER TABLE public.todos
ADD CONSTRAINT todos_users
FOREIGN KEY (user_id)
REFERENCES public.users (id)
ON DELETE CASCADE;