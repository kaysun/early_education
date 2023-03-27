CREATE TABLE `user` (
    `user_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户id',
    `name` varchar(128) NOT NULL DEFAULT '' COMMENT '名字',
    `password` varchar(128) NOT NULL DEFAULT '' COMMENT '密码，密文',
    `nick` varchar(128) NOT NULL DEFAULT '' COMMENT '昵称',
    `birthday` datetime NOT NULL DEFAULT '0' COMMENT '生日',
    `phone` int(11) NOT NULL DEFAULT 0 COMMENT '手机号',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`user_id`),
    Index `index_name_passwd` (`name`, `password`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户表';

CREATE TABLE `course` (
    `course_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '课程id',
    `name` varchar(128) NOT NULL DEFAULT '' COMMENT '课程名字',
    `class_time` datetime NOT NULL COMMENT '上课时间',
    `teacher_id` int(11) NOT NULL COMMENT '教师id',
    `max_students` tinyint NOT NULL DEFAULT 0 COMMENT '课程最多容纳的学生数量',
    `scheduled_students` tinyint NOT NULL DEFAULT 0 COMMENT '已经约课的学生数',
    `course_status` tinyint NOT NULL DEFAULT 0 COMMENT '课程状态：1-有空闲 2-约满',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`course_id`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '课程表';



