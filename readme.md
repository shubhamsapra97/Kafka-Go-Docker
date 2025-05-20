1. Setup postgres in local
    i. brew install postgresql
    ii. brew services start postgresql
    iii. psql postgres
    iv. CREATE USER user WITH PASSWORD 'password';
    v. CREATE DATABASE taskdb OWNER user;
    vi. psql -U user -d taskdb -h localhost
