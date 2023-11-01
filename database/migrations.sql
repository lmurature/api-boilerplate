CREATE DATABASE users_db /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

use users_db;

-- users_db.`user` definition

CREATE TABLE `user` (
                        `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                        `first_name` varchar(100) DEFAULT NULL,
                        `last_name` varchar(100) DEFAULT NULL,
                        `email` varchar(100) DEFAULT NULL,
                        `password` varchar(512) DEFAULT NULL,
                        `contact_phone` varchar(32) DEFAULT NULL,
                        `user_type` varchar(100) DEFAULT NULL,
                        `date_created` datetime DEFAULT NULL,
                        `refresh_token` varchar(1024) DEFAULT NULL,
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;