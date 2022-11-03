create table users(
      id serial primary key,
      first_name varchar,
      last_name varchar,
      email varchar,
      mobile_number varchar,
      password varchar,
      birthday int,
      gender varchar,
      active int,
      path_account_activation varchar,
      created_at int,
      updated_at int,
      deleted_at int
);

create table password_reset_configs(
       id serial primary key,
       users_id int,
       url varchar,
       expiration_by_use int,
       updated_at int,
       foreign key(users_id) references users(id)
);

create table friendships(
        id serial primary key,
        requester_id int,
        requested_id int,
        requested_at int,
        accepted_at int,
        accepted int,
        foreign key(requester_id) references users(id),
        foreign key(requested_id) references users(id)
);
