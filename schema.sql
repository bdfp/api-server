CREATE DATABASE IF NOT EXISTS shadowfax;

DROP TABLE IF EXISTS business_review;
DROP TABLE IF EXISTS business_service;
DROP TABLE IF EXISTS business_tags;
DROP TABLE IF EXISTS service;
DROP TABLE IF EXISTS tag;
DROP TABLE IF EXISTS review;
DROP TABLE IF EXISTS user;
DROP TABLE IF EXISTS business;

CREATE TABLE business (
  id int(11) AUTO_INCREMENT PRIMARY KEY,
  name varchar(256) NOT NULL,
  city varchar(256) NOT NULL,
  primary_email varchar(256) NOT NULL,
  primary_phone varchar(256) NOT NULL,
  latitude varchar(10) DEFAULT NULL,
  longitude varchar(10) DEFAULT NULL,
  overall_rating int(11) DEFAULT NULL
);

--
-- Table structure for table user
--

CREATE TABLE user (
  id int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name varchar(256) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Table structure for table reviews
--

CREATE TABLE review (
  id int(11) AUTO_INCREMENT PRIMARY KEY,
  rating int(11) DEFAULT NULL,
  comment varchar(256) DEFAULT NULL,
  user_id int(11) DEFAULT NULL,
  FOREIGN KEY (user_id) REFERENCES user(id)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------


--
-- Table structure for table business_review
--

CREATE TABLE business_review (
  id int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  review_id int(11) DEFAULT NULL,
  user_id int(11) DEFAULT NULL,
  FOREIGN KEY (review_id) REFERENCES review(id),
  FOREIGN KEY (user_id) REFERENCES user(id)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table services
--

CREATE TABLE service (
  id int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name varchar(256) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table business_service
--

CREATE TABLE business_service (
  id int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  service_id int(11) DEFAULT NULL ,
  business_id int(11) DEFAULT NULL ,
  FOREIGN KEY (business_id) REFERENCES business(id),
  FOREIGN KEY (service_id) REFERENCES service(id)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table tags
--

CREATE TABLE tag (
  id int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name varchar(256) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table business_tags
--

CREATE TABLE business_tags (
  id int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  tag_id int(11) DEFAULT NULL ,
  business_id int(11) DEFAULT NULL ,
  FOREIGN KEY (business_id) REFERENCES business(id),
  FOREIGN KEY (tag_id) REFERENCES tag(id)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------






