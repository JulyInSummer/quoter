create table if not exists quotes(
    id serial primary key,
    author varchar(50),
    quote varchar(500)
);