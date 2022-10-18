-- create new database
CREATE DATABASE `freefolks-fc`;

-- player_information
CREATE TABLE player_information(
  pi_id int NOT NULL AUTO_INCREMENT,
  pi_user_type tinyint,
  pi_customer_number varchar(255),
  pi_first_name varchar(255),
  pi_last_name varchar(255),
  pi_email varchar(255),
  pi_password varchar(255),
  pi_mobile_number varchar(255),
  pi_occupation int COMMENT 'Reference to mr_occupation',
  pi_date_of_birth date,
  pi_gender varchar(255),
  pi_photo_profile varchar(255),
  pi_address text,
  pi_city varchar(255),
  pi_postal_code varchar(255),
  pi_body_size enum('XXS', 'XS', 'S', 'M', 'L', 'XL', 'XXL', 'XXXL'),
  pi_activation_code varchar(255),
  pi_email_status enum('Y', 'N'),
  pi_verified_at DATETIME,
  pi_created_at TIMESTAMP DEFAULT NOW(),
  pi_created_by int,
  pi_updated_at TIMESTAMP DEFAULT NOW() ON UPDATE NOW(),
  pi_updated_by int,
  PRIMARY KEY (pi_id)
);

-- preferred_position
CREATE TABLE preferred_position(
  pp_id int NOT NULL AUTO_INCREMENT,
  pp_pi_id int,
  pp_position int COMMENT 'Reference to mr_position',
  pp_created_at TIMESTAMP DEFAULT NOW(),
  pp_created_by int,
  pp_updated_at TIMESTAMP DEFAULT NOW() ON UPDATE NOW(),
  pp_updated_by int,
  PRIMARY KEY (pp_id),
  FOREIGN KEY (pp_pi_id) REFERENCES player_information(pi_id)
);

-- game_data
CREATE TABLE game_data(
  gd_id int NOT NULL AUTO_INCREMENT,
  gd_game_number varchar(255),
  gd_venue_name varchar(255),
  gd_venue_address varchar(255),
  gd_map_url varchar(255),
  gd_time DATETIME,
  gd_goalkeeper_quota int,
  gd_outfield_quota int,
  gd_goalkeeper_price double,
  gd_outfield_price double,
  gd_total_cost double,
  gd_notes text,
  gd_status enum('Y', 'N'),
  gd_created_at TIMESTAMP DEFAULT NOW(),
  gd_created_by int,
  gd_updated_at TIMESTAMP DEFAULT NOW() ON UPDATE NOW(),
  gd_updated_by int,
  PRIMARY KEY (gd_id)
);

-- game_information
CREATE TABLE game_information(
  gi_id int NOT NULL AUTO_INCREMENT,
  gi_gd_id int,
  gi_type int,
  gi_description text,
  gi_created_at TIMESTAMP DEFAULT NOW(),
  gi_created_by int,
  gi_updated_at TIMESTAMP DEFAULT NOW() ON UPDATE NOW(),
  gi_updated_by int,
  PRIMARY KEY (gi_id),
  FOREIGN KEY (gi_gd_id) REFERENCES game_data(gd_id)
);

-- game_galleries
CREATE TABLE game_galleries(
  ggs_id int NOT NULL AUTO_INCREMENT,
  ggs_gd_id int,
  ggs_image_url text,
  ggs_alt_image varchar(255),
  ggs_created_at TIMESTAMP DEFAULT NOW(),
  ggs_created_by int,
  ggs_updated_at TIMESTAMP DEFAULT NOW() ON UPDATE NOW(),
  ggs_updated_by int,
  PRIMARY KEY (ggs_id),
  FOREIGN KEY (ggs_gd_id) REFERENCES game_data(gd_id)
)

-- game_costs
CREATE TABLE game_costs(
  gcs_id int NOT NULL AUTO_INCREMENT,
  gcs_gd_id int,
  gcs_description text,
  gcs_cost double,
  gcs_created_at TIMESTAMP DEFAULT NOW(),
  gcs_created_by int,
  gcs_updated_at TIMESTAMP DEFAULT NOW() ON UPDATE NOW(),
  gcs_updated_by int,
  PRIMARY KEY (gcs_id)
  FOREIGN KEY (gcs_gd_id) REFERENCES game_data(gd_id)
)

-- game_registration
CREATE TABLE game_registration(
  gr_id int NOT NULL AUTO_INCREMENT,
  gr_gd_id int,
  gr_pi_id int,
  gr_is_outfield enum('Y', 'N'),
  gr_amount double,
  gr_transaction_number varchar(255),
  gr_created_at TIMESTAMP DEFAULT NOW(),
  gr_created_by int,
  gr_updated_at TIMESTAMP DEFAULT NOW() ON UPDATE NOW(),
  gr_updated_by int,
  PRIMARY KEY (gr_id),
  FOREIGN KEY (gr_gd_id) REFERENCES game_data(gd_id),
  FOREIGN KEY (gr_pi_id) REFERENCES player_information(pi_id)
)

-- game_registered_player
CREATE TABLE game_registered_player(
  grp_id int NOT NULL AUTO_INCREMENT,
  grp_gd_id int,
  grp_pi_id int,
  grp_is_outfield enum('Y', 'N'),
  grp_amount_paid double,
  grp_paid_at DATETIME,
  grp_transaction_number varchar(255),
  grp_created_at TIMESTAMP DEFAULT NOW(),
  grp_created_by int,
  grp_updated_at TIMESTAMP DEFAULT NOW() ON UPDATE NOW(),
  grp_updated_by int,
  PRIMARY KEY (grp_id),
  FOREIGN KEY (grp_gd_id) REFERENCES game_data(gd_id),
  FOREIGN KEY (grp_pi_id) REFERENCES player_information(pi_id)
)

-- 