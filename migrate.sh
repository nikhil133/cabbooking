#!/bin/bash
psql postgres -c "drop database if exists cabbooking;"
psql postgres -c "drop user if exists cab;"
psql postgres -c "create database cabbooking;"
psql postgres -c "create user cab with encrypted password 'cabpassword';"
psql postgres -c "grant all privileges on database cabbooking to cab;"

psql cabbooking < cab_sql_dump