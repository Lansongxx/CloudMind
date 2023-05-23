create table file
(
    id         bigint auto_increment
        primary key,
    created_at datetime     null,
    updated_at datetime     null,
    deleted_at datetime     null,
    name       varchar(50)  not null,
    type       varchar(20)  not null,
    path       varchar(100) not null,
    size       varchar(20)  not null,
    share_link varchar(100) not null,
    parent_id  bigint       not null,
    constraint file_pk2
        unique (id)
);

create table file_md5
(
    id      bigint       not null
        primary key,
    file_id bigint       not null,
    md5     varchar(500) not null,
    constraint file_md5_pk2
        unique (md5)
);

create table profile_picture
(
    id      bigint auto_increment
        primary key,
    md5     varchar(600) not null,
    picture mediumblob   not null,
    constraint Profile_Picture_pk
        unique (md5)
);

create table user_file
(
    id      bigint auto_increment
        primary key,
    user_id bigint not null,
    file_id bigint not null,
    constraint user_file_pk
        unique (file_id),
    constraint user_file_pk2
        unique (id)
);

create table user_profile
(
    id      bigint auto_increment
        primary key,
    user_id bigint       not null,
    md5     varchar(600) null
);



// 使用joins 生成衍生表(等价于子查询)
query := db.Table("order").Select("MAX(order.finished_at) as latest").Joins("left join user user on order.user_id = user.id").Where("user.age > ?", 18).Group("order.user_id")
db.Model(&Order{}).Joins("join (?) q on order.finished_at = q.latest", query).Scan(&results)
原生SQL
SELECT `order`.`user_id`, `order`.`finished_at` -- 选择查询结果需要返回的字段
FROM `order` -- 操作对象是 order 表
JOIN -- 使用 JOIN 关键字连接 order 表和子查询 q
(
    SELECT MAX(order.finished_at) AS latest -- 计算每个用户最近一次订单的完成时间
    FROM `order`
    LEFT JOIN user -- 对 order 表和 user 表使用 LEFT JOIN 进行关联
    ON order.user_id = user.id
    WHERE user.age > 18 -- 筛选年龄大于 18 岁的用户
    GROUP BY `order`.`user_id` -- 按照用户 ID 进行分组
) q ON order.finished_at = q.latest -- 连接条件为订单的完成时间等于该用户最近一次订单的完成时间

