CREATE DATABASE IF NOT EXISTS rsql;
USE rsql;
CREATE TABLE IF NOT EXISTS `posts` (
  `post_id` int(11) NOT NULL AUTO_INCREMENT,
  `post_title` varchar(100) NOT NULL,
  `post_body` text NOT NULL,
  PRIMARY KEY (`post_id`)
);