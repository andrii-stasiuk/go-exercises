CREATE TABLE public.states (
    id bigserial,
    state text NOT NULL UNIQUE
);



INSERT INTO `public.states` (`id`, `state`) VALUES
(1, 'created'),
(2, 'wait'),
(3, 'canceled'),
(4, 'blocked'),
(5, 'in process/doing'),
(6, 'review'),
(7, 'done'),
(8, 'archived');