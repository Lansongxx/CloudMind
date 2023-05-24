create table feedback
(
    feedback_type varchar(256) not null,
    user_id       varchar(256) not null,
    item_id       varchar(256) not null,
    time_stamp    timestamp    not null,
    comment       text         not null,
    primary key (feedback_type, user_id, item_id)
)
    charset = utf8mb4;

create index item_id
    on feedback (item_id);

create index user_id
    on feedback (user_id);

create table items
(
    item_id    varchar(256) not null
        primary key,
    time_stamp timestamp    not null,
    labels     json         not null,
    comment    text         not null,
    is_hidden  tinyint(1)   not null,
    categories json         not null
)
    charset = utf8mb4;

create table users
(
    user_id   varchar(256) not null
        primary key,
    labels    json         not null,
    comment   text         not null,
    subscribe json         not null
)
    charset = utf8mb4;

