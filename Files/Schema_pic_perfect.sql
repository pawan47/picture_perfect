CREATE table movie_info(
  movie_id SERIAL,
  movie_name varchar not NULL,
  short_discription varchar,
  long_discription varchar,
  thumbnail_link varchar,
  genre varchar,
  language varchar,
  actor varchar,
  actress varchar,
  director varchar,
  PRIMARY key(movie_id)
  );
  
CREATE TABLE user_details(
  user_id SERIAL,
  first_name varchar(50) not NULL,
  last_name VARCHAR(50),
  username VARCHAR(50) not NULL UNIQUE,
  password VARCHAR NOT NULL,
  is_admin boolean,
  PRIMARY kEY(user_id)
  );
  
 CREATE TABLE rating_review(
   movie_id SERIAL,
   user_id SERIAL,
   rating integer,
   review varchar,
   FOREIGN key (movie_id) REFERENCES movie_info(movie_id),
   foreign key (user_id) REFERENCES user_details(user_id),
   PRIMARY key(movie_id,user_id)
   );

 
 CREATE table cineplex_info(
   cineplex_id SERIAL,
   cineplex_name varchar not NULL,
   city Varchar not NULL,
   PRIMARY key (cineplex_id)
  );

 create table show_id(
   show_id SERIAL,
   movie_id SERIAL,
   cineplex_id SERIAL,
   starting_time TIME not NULL,
   end_time TIME not NULL,
   PRIMARY key (show_id),
   FOREIGN key (cineplex_id) REFERENCES cineplex_info(cineplex_id),
   FOREIGN key(movie_id) REFERENCES movie_info(movie_id)
  );
