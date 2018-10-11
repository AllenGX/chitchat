drop table users;

create table users (
  userid     int(4) auto_increment primary key not null,
  username   varchar(16) not null,
  email      varchar(32) not null unique,
  password   varchar(50) not null
);
