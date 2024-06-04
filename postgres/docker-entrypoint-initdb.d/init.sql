CREATE USER owner;

CREATE DATABASE "auth-service";
GRANT ALL PRIVILEGES ON DATABASE "auth-service" TO owner;
GRANT ALL PRIVILEGES ON DATABASE "auth-service" TO root;

CREATE DATABASE "registration-service";
GRANT ALL PRIVILEGES ON DATABASE "registration-service" TO owner;
GRANT ALL PRIVILEGES ON DATABASE "registration-service" TO root;

CREATE DATABASE "user-service";
GRANT ALL PRIVILEGES ON DATABASE "user-service" TO owner;
GRANT ALL PRIVILEGES ON DATABASE "user-service" TO root;