Prerequisites:
Postgres
goose


How to install everything on linux:
Postgres:

    sudo apt update
    sudo apt install postgresql postgresql-contrib

check it worked by running:

    psql --version

It should say you are on version 14+

Set the system password for the postgres user by running:
    
    sudo passwd postgres

Dont forget this password

Start the server in the background
    
    sudo service postgresql start

Enter the psql shell:

    sudo -u postgres psql

Create a new database:

    CREATE DATABASE pressure;

Connect to the database 

    \c pressure

set the postgres user password

    ALTER USER postgres PASSWORD 'a password';
