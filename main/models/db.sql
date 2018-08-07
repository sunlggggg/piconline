create table user
(
  id            bigint auto_increment
    primary key,
  name          varchar(100) not null,
  email         varchar(40)  not null,
  register_time bigint       not null
);

create table picture
(
  id        bigint auto_increment
    primary key,
  name      varchar(20)            not null,
  is_public tinyint(1) default '0' not null,
  time      bigint                 not null,
  `desc`    varchar(50)            null,
  user_id   bigint                 null,
  constraint picture_id_uindex
  unique (id),
  constraint picture_name_uindex
  unique (name),
  constraint picture_user_id_fk
  foreign key (user_id) references user (id)
);

create index picture_user_id_fk
  on picture (user_id);
