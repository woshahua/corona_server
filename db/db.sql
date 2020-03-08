create table patient_news (
  id serial not null,
  title varchar(255) not null,
  link text not null,
  primary key (id)
);

create table patient_new_patient (
    id serial not nll,
    sum integer not null,
    growth integer not null,
    created_on timestamp with time zone,
    primary key (id)
);

create table patient_dead_patient (
    id serial not nll,
    sum integer not null,
    growth integer not null,
    created_on time not null,
    primary key (id)
);

create table patient_japanese (
    id integer not null,
    date varchar(20) not null,
    age varchar(10) not null,
    patient_location varchar(30) not null,
    primary key (id)
);