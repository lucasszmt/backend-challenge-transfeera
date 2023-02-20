package main

const (
	CreateExtension       = `CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`
	CreateTablePixKeyType = `CREATE TABLE IF NOT EXISTS pix_key_type
	(
		id   uuid PRIMARY KEY NOT NULL default uuid_generate_v4(),
		name varchar(50) UNIQUE
	)`
	InsertPixKeyItems = `INSERT INTO pix_key_type(name)
	VALUES ('cpf'),
		   ('cnpj'),
		   ('email'),
		   ('phone'),
		   ('random_key')`
	CreateReceiverTable = `CREATE TABLE IF NOT EXISTS receiver
	(
		id         uuid PRIMARY KEY NOT NULL,
		name       varchar(255),
		email      varchar(250),
		document   varchar(15),
		pixKey     varchar(50),
		pixKeyType uuid references pix_key_type (id),
		status     int
	)`
)
