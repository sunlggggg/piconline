create table user
(
  id            bigint auto_increment
    primary key,
  name          varchar(100) not null,
  email         varchar(40)  not null,
  register_time bigint       not null,
  password      varchar(20)  null,
  constraint user_name_uindex
  unique (name)
);

create table content
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
  on content (user_id);

create table file
(
  id         bigint auto_increment
    primary key,
  isFile     tinyint(1) default '0' null,
  leftChild  bigint                 null,
  rightChild bigint                 null,
  contentID  bigint                 null,
  createTime bigint                 not null,
  updateTime bigint                 not null,
  constraint file_id_uindex
  unique (id),
  constraint file_content_id_fk
  foreign key (contentID) references content (id)
);

create index file_content_id_fk
  on file (contentID);

create table fileroot
(
  id         bigint auto_increment
    primary key,
  userId     bigint null,
  fileId     bigint null,
  createTime int    null,
  constraint table_name_id_uindex
  unique (id),
  constraint table_name_user_id_fk
  foreign key (userId) references user (id),
  constraint fileroot_file_id_fk
  foreign key (fileId) references file (id)
);

create index fileroot_file_id_fk
  on fileroot (fileId);

create index table_name_user_id_fk
  on fileroot (userId);

