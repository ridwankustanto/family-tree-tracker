CREATE TABLE IF NOT EXISTS family_origins (
  id CHAR(32) PRIMARY KEY,
  name VARCHAR(64) NOT NULL,
  alias VARCHAR(1) NOT NULL,
  description VARCHAR(255) NOT NULL,
  address VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS members (
  id CHAR(32) PRIMARY KEY,
  family_id VARCHAR(64) NOT NULL,
  name VARCHAR(64) NOT NULL,
  nickname VARCHAR(16) NOT NULL,
  place_of_birth VARCHAR(64) NOT NULL,
  date_of_birth VARCHAR(64) NOT NULL,
  origin_address VARCHAR(255) NOT NULL,
  current_address VARCHAR(255) NOT NULL,
  mobile_phone VARCHAR(16) NOT NULL,
  phone VARCHAR(16) NOT NULL,
  religion VARCHAR(16) NOT NULL,
  photo VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS member_has_spouse (
  id CHAR(32) PRIMARY KEY,
  member_id CHAR(32) NOT NULL,
  name VARCHAR(64) NOT NULL,
  nickname VARCHAR(16) NOT NULL,
  place_of_birth VARCHAR(64) NOT NULL,
  date_of_birth VARCHAR(64) NOT NULL,
  origin_address VARCHAR(255) NOT NULL,
  photo VARCHAR(255) NOT NULL,
  CONSTRAINT FK_SpouseMember FOREIGN KEY (member_id)
  REFERENCES members(member_id)
);

CREATE TABLE IF NOT EXISTS member_has_childs (
  id CHAR(32) PRIMARY KEY,
  member_id CHAR(32) NOT NULL,
  family_id VARCHAR(64) NOT NULL,
  name VARCHAR(64) NOT NULL,
  nickname VARCHAR(16) NOT NULL,
  place_of_birth VARCHAR(64) NOT NULL,
  date_of_birth VARCHAR(64) NOT NULL,
  origin_address VARCHAR(255) NOT NULL,
  gender VARCHAR(1) NOT NULL,
  photo VARCHAR(255) NOT NULL,
  CONSTRAINT FK_ChildMember FOREIGN KEY (member_id)
  REFERENCES members(member_id)
);
