Prerequisites:
golang
Postgres
goose



How to install everything on linux:
Golang:  go to https://www.go.dev/doc/install for up to date instructions.
These instructions were written nov 18 2024

Download archive:

    https://go.dev/dl/go1.23.3.linux-amd64.tar.gz


Remove any previous Go installation by deleting the /usr/local/go folder (if it exists), then extract the archive you just downloaded into /usr/local, creating a fresh Go tree in /usr/local/go:

      $ rm -rf /usr/local/go && tar -C /usr/local -xzf go1.23.3.linux-amd64.tar.gz
      

(You may need to run the command as root or through sudo).

Do not untar the archive into an existing /usr/local/go tree. This is known to produce broken Go installations.
Add /usr/local/go/bin to the PATH environment variable.

You can do this by adding the following line to your $HOME/.profile or /etc/profile (for a system-wide installation):

          export PATH=$PATH:/usr/local/go/bin
          

Note: Changes made to a profile file may not apply until the next time you log into your computer. To apply the changes immediately, just run the shell commands directly or execute them from the profile using a command such as source $HOME/.profile.
Verify that you've installed Go by opening a command prompt and typing the following command:

          $ go version
          

Confirm that the command prints the installed version of Go.





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

Goose:
Install goose with the go package manager

    go install github.com/pressly/goose/v3/cmd/goose@latest



Navigate to the /sql/schema folder and run a goose up command.
It will look like this but you will need to fill in your own postgres username
and password
    
    goose postgres postgres://username here:password here@localhost:5432/pressure?sslmode=disable up
