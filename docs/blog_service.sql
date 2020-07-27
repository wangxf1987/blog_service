drop database if exists blog_service;
create database blog_service;
use blog_service;

drop table if exists blog_tag;
drop table if exists blog_article;

/*==============================================================*/
/* Table: blog_tag                                               */
/*==============================================================*/
create table blog_tag
(
   id                   int(10) unsigned not null auto_increment,
   name                 varchar(100) default '' commit '标签名称',
   state                tinyint(3) unsigned default '1' commit '状态 0 为禁用，1 为启用',
   create_on            int(10) unsigned default '0' commit '创建时间',
   create_by            varchar(100) default '' commit '创建人',
   modified_on          int(10) unsigned default '0' commit '修改时间',
   modified_by          varchar(100) default '' commit '修改人',
   delete_at            int(10) unsigned default '0' commit '删除时间',
   is_delete            tinyint(3) unsigned default '0' commit '是否删除 0 为未删除, 1 为已删除',
   primary key (id)
)
DEFAULT charset = UTF8
ENGINE = InnoDB
COLLATE = utf8_general_ci;

alter table actions comment '标签管理';

/*==============================================================*/
/* Table: blog_article                                          */
/*==============================================================*/
create table blog_article
(
   id                   int(10) unsigned not null auto_increment,
   title                varchar(100) default '' commit '文章标题',
   desc                 varchar(255) default '' commit '文章描述',
   cover_image_url      varchar(255) default '' commit '封面图片地址',
   content              longtext commit '文章内容',
   create_on            int(10) unsigned default '0' commit '创建时间',
   create_by            varchar(100) default '' commit '创建人',
   modified_on          int(10) unsigned default '0' commit '修改时间',
   modified_by          varchar(100) default '' commit '修改人',
   delete_at            int(10) unsigned default '0' commit '删除时间',
   is_delete            tinyint(3) unsigned default '0' commit '是否删除 0 为未删除, 1 为已删除',
   primary key (id)
)
DEFAULT charset = UTF8
ENGINE = InnoDB
COLLATE = utf8_general_ci;

alter table actions comment '文章管理';

/*==============================================================*/
/* Table: blog_article_tag                                          */
/*==============================================================*/
create table blog_article_tag
(
   id                   int(10) unsigned not null auto_increment,
   article_id           int(11) not null commit '文章ID',
   tag_id               int(10) unsigned not null default '0' commit '标签ID',
   create_on            int(10) unsigned default '0' commit '创建时间',
   create_by            varchar(100) default '' commit '创建人',
   modified_on          int(10) unsigned default '0' commit '修改时间',
   modified_by          varchar(100) default '' commit '修改人',
   delete_at            int(10) unsigned default '0' commit '删除时间',
   is_delete            tinyint(3) unsigned default '0' commit '是否删除 0 为未删除, 1 为已删除',
   primary key (id)
)
DEFAULT charset = UTF8
ENGINE = InnoDB
COLLATE = utf8_general_ci;

alter table actions comment '文章标签关联';