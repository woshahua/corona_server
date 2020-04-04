create table `patient_locations` (
    `id` int(10) NOT NULL AUTO_INCREMENT,
    `location` varchar(10) NOT NULL UNIQUE,
    `sum` int(10) NOT NULL,
    `created_at` datetime NOT NULL,
    `updated_at` datetime NOT NULL,
    `deleted_at` datetime,
    primary key (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

create table `patient_by_dates` (
    `id` int(10) NOT NULL AUTO_INCREMENT,
    `date` varchar(10) NOT NULL UNIQUE,
    `confirmed` int(10) NOT NULL,
    `recovered` int(10) NOT NULL,
    `dead` int(10) NOT NULL,
    `critical` int(10) NOT NULL,
    `tested` int(10) NOT NULL,
    `created_at` datetime NOT NULL,
    `updated_at` datetime NOT NULL,
    `deleted_at` datetime,
    primary key (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
