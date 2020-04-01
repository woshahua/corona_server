drop table if exists patient_location;
create table `patient_location` (
    `id` int(10) NOT NULL AUTO_INCREMENT,
    `location` varchar(10) NOT NULL UNIQUE,
    `sum` int(10) NOT NULL,
    primary key (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

drop table if exists patient_location;
create table `patient_by_date` (
    `id` int(10) NOT NULL AUTO_INCREMENT,
    `date` varchar(10) NOT NULL UNIQUE,
    `confirmed` int(10) NOT NULL,
    `recovered` int(10) NOT NULL,
    `dead` int(10) NOT NULL,
    `critical` int(10) NOT NULL,
    `tested` int(10) NOT NULL,
    primary key (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

drop table if exists news;
create table `news` (
    `id` int(10) NOT NULL AUTO_INCREMENT,
    `title` varchar(10) NOT NULL UNIQUE,
    `link` varchar(10) NOT NULL UNIQUE,
    `url` varchar(10) NOT NULL UNIQUE,
    `updated_time` varchar(10) NOT NULL,
    `passed_day` int(10) NOT NULL,
    `passed_minutes` int(10) NOT NULL,
    `passed_hour` int(10) NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    primary key (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

drop table if exists last_update_time;
create table `last_update_time`(
    `id` int(10) NOT NULL AUTO_INCREMENT,
    `patient_data_update_time` varchar(10) NOT NULL,
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
