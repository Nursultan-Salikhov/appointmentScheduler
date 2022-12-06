CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE schedules
(
    id serial not null unique,
    user_id int references users (id) on delete cascade not null,
    work_day date not null,
    start_time time not null,
    end_time time not null
);

CREATE TABLE clients
(
    id serial not null unique,
    user_id int references users (id) on delete cascade not null,
    name varchar(255) not null,
    phone_number varchar(20),
    email varchar(255),
    tg_username varchar(255),
    description varchar(255)
);

CREATE TABLE appointments
(
    id serial not null unique,
    appointment_day date not null,
    appointment_time time not null
);

CREATE TABLE clients_appointments
(
    id serial not null unique,
    client_id int references clients (id) on delete cascade not null,
    appointment_id int references appointments (id) on delete cascade not null
);

CREATE TABLE notices_templates
(
    id serial not null unique,
    user_id int references users (id) on delete cascade not null unique,
    appointment_template varchar(255),
    reminder_template varchar(255)
);

CREATE TABLE email_settings
(
    id serial not null unique,
    user_id int references users (id) on delete cascade not null,
    status boolean not null default false,
    email varchar(255),
    password_hash varchar(255),
    host varchar(255),
    port int
);