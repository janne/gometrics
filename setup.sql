create table metrics
(
  id serial primary key not null,
  time timestamp not null,
  key varchar(255) not null,
  value float not null
)
