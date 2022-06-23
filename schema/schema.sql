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
  province VARCHAR(32) NOT NULL,
  city VARCHAR(32) NOT NULL,
  disctrict VARCHAR(32) NOT NULL,
  subdisctrict VARCHAR(32) NOT NULL,
  mobile_phone VARCHAR(16) NOT NULL,
  phone VARCHAR(16) NOT NULL,
  religion VARCHAR(16) NOT NULL,
  gender VARCHAR(1) NOT NULL,
  photo VARCHAR(255) NOT NULL,
  is_origin BOOLEAN NOT NULL,
  created_at DATE NOT NULL,
  updated_at DATE NOT NULL
);

CREATE TABLE IF NOT EXISTS accounts (
  id CHAR(32) PRIMARY KEY,
  people_id CHAR(32) NOT NULL,
  username VARCHAR(64) NOT NULL,
  password VARCHAR(128) NOT NULL,
  created_at DATE NOT NULL,
  updated_at DATE NOT NULL,
  is_admin BOOLEAN NOT NULL,
  CONSTRAINT FK_AccountPeople FOREIGN KEY (people_id)
  REFERENCES people(id)
);

CREATE TABLE IF NOT EXISTS people_has_relations (
  id CHAR(32) PRIMARY KEY,
  people_id CHAR(32) NOT NULL,
  with_people_id CHAR(32) NOT NULL,
  relation_type VARCHAR(8) NOT NULL,
  created_at DATE NOT NULL,
  updated_at DATE NOT NULL,
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
  created_at DATE NOT NULL,
  updated_at DATE NOT NULL,
  CONSTRAINT FK_FamilyHusbandPeople FOREIGN KEY (husband_id)
  REFERENCES people(id),
  CONSTRAINT FK_EventWifePeople FOREIGN KEY (wife_id)
  REFERENCES people(id)
);


CREATE TABLE IF NOT EXISTS people_has_events (
  id CHAR(32) PRIMARY KEY,
  people_id CHAR(32) NOT NULL,
  category VARCHAR(16) NOT NULL,
  date DATE NOT NULL,
  created_at DATE NOT NULL,
  updated_at DATE NOT NULL,
  CONSTRAINT FK_EventPeople FOREIGN KEY (people_id)
  REFERENCES people(id)
);

CREATE TABLE IF NOT EXISTS log_activities (
  id CHAR(32) PRIMARY KEY,
  people_id CHAR(32) NOT NULL,
  action VARCHAR(255) NOT NULL,
  ip_address VARCHAR(16) NOT NULL,
  client VARCHAR(65) NOT NULL,
  created_at DATE NOT NULL,
  updated_at DATE NOT NULL,
  CONSTRAINT FK_LogPeople FOREIGN KEY (people_id)
  REFERENCES people(id)
);

CREATE TABLE IF NOT EXISTS form_verifications (
  id CHAR(32) PRIMARY KEY,
  generated_by CHAR(32) NOT NULL,
  used_by CHAR(32),
  code VARCHAR(64) NOT NULL,
  is_used BOOLEAN NOT NULL,
  created_at DATE NOT NULL,
  updated_at DATE NOT NULL,
  CONSTRAINT FK_FormGeneratedBy FOREIGN KEY (generated_by)
  REFERENCES people(id),
  CONSTRAINT FK_FormUsedBy FOREIGN KEY (used_by)
  REFERENCES people(id)
);

CREATE TABLE IF NOT EXISTS country (
  id CHAR(32) PRIMARY KEY,
  name VARCHAR(64) NOT NULL,
  code VARCHAR(64) NOT NULL,
  created_at DATE NOT NULL,
  updated_at DATE NOT NULL,
);

CREATE TABLE IF NOT EXISTS provinces (
  id CHAR(32) PRIMARY KEY,
  country_id CHAR(32) NOT NULL,
  name VARCHAR(64) NOT NULL,
  code VARCHAR(64) NOT NULL,
  created_at DATE NOT NULL,
  updated_at DATE NOT NULL,
  CONSTRAINT FK_Country FOREIGN KEY (country_id)
  REFERENCES country(id)
);

CREATE TABLE IF NOT EXISTS city (
  id CHAR(32) PRIMARY KEY,
  province_id CHAR(32) NOT NULL,
  name VARCHAR(64) NOT NULL,
  code VARCHAR(64) NOT NULL,
  created_at DATE NOT NULL,
  updated_at DATE NOT NULL,
  CONSTRAINT FK_Provinces FOREIGN KEY (province_id)
  REFERENCES provinces(id)
);

CREATE TABLE IF NOT EXISTS districts(
  id CHAR(32) PRIMARY KEY,
  city_id CHAR(32) NOT NULL,
  name VARCHAR(64) NOT NULL,
  code VARCHAR(64) NOT NULL,
  created_at DATE NOT NULL,
  updated_at DATE NOT NULL,
  CONSTRAINT FK_City FOREIGN KEY (city_id)
  REFERENCES city(id)
);

CREATE TABLE IF NOT EXISTS subdistricts(
  id CHAR(32) PRIMARY KEY,
  district_id CHAR(32) NOT NULL,
  name VARCHAR(64) NOT NULL,
  code VARCHAR(64) NOT NULL,
  created_at DATE NOT NULL,
  updated_at DATE NOT NULL,
  CONSTRAINT FK_City FOREIGN KEY (district_id)
  REFERENCES districts(id)
);