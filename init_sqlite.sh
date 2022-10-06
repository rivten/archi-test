#!/usr/bin/sh


sqlite3 db.db < sql/init.sql 
sqlite3 db.db < sql/fixtures.sql
