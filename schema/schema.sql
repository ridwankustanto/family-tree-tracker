CREATE TABLE IF NOT EXISTS accounts (
  id CHAR(32) PRIMARY KEY,
  username VARCHAR(64) NOT NULL,
  password VARCHAR(1) NOT NULL,
  is_admin BOOLEAN NOT NULL
);

CREATE TABLE IF NOT EXISTS people (
  id CHAR(32) PRIMARY KEY,
  family_id VARCHAR(64) NOT NULL,
  fullname VARCHAR(64) NOT NULL,
  nickname VARCHAR(16) NOT NULL,
  place_of_birth VARCHAR(64) NOT NULL,
  date_of_birth VARCHAR(64) NOT NULL,
  origin_address VARCHAR(255) NOT NULL,
  current_address VARCHAR(255) NOT NULL,
  country VARCHAR(32) NOT NULL,
  state VARCHAR(32) NOT NULL,
  city VARCHAR(32) NOT NULL,
  disctrict VARCHAR(32) NOT NULL,
  subdisctrict VARCHAR(32) NOT NULL,
  mobile_phone VARCHAR(16) NOT NULL,
  phone VARCHAR(16) NOT NULL,
  religion VARCHAR(16) NOT NULL,
  gender VARCHAR(1) NOT NULL,
  photo VARCHAR(255) NOT NULL,
  is_origin BOOLEAN NOT NULL
);

CREATE TABLE IF NOT EXISTS people_has_relations (
  id CHAR(32) PRIMARY KEY,
  people_id CHAR(32) NOT NULL,
  with_people_id CHAR(32) NOT NULL,
  relation_type VARCHAR(8) NOT NULL,
  CONSTRAINT FK_RelationPeople FOREIGN KEY (people_id)
  REFERENCES people(id),
  CONSTRAINT FK_RelationWithPeople FOREIGN KEY (with_people_id)
  REFERENCES people(id)
);

CREATE TABLE IF NOT EXISTS people_has_family (
  id CHAR(32) PRIMARY KEY,
  family_name VARCHAR(64) NOT NULL,
  husband_id CHAR(32) NOT NULL,
  wife_id CHAR(32),
  CONSTRAINT FK_FamilyPeople FOREIGN KEY (husband_id)
  REFERENCES people(id)
);
