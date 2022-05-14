#Database:

Create a migration using [dmate](https://github.com/amacneil/dbmate). Dbmate uses .env's dburl in order to connect to the database

Dbmate commands:

    dbmate init [name of database]
    dbmate up
    dbmate down

- dbmate init creates database with name of file
- dbmate up runs migration updates tables 
- dbmate down runs migration and deletes tables

Dbmate up runs once. if you make edits to the sql file you have to run dbmate down in order for the database to be updated.

##Test data:
[mockaroo](https://www.mockaroo.com/) is used to create sample random data. You can add your table field names and convert to sql file in order to insert to your database

If you create a test files to fill your database

On windows:
if psql in your path follow mac instructions. if not use .bat file in order to automate the process. 
```
REM Run psql
[path to .exe psql]"C:\Program Files\PostgreSQL\14\bin\psql.exe" -h "[hostname]" -U "[username]" -d "[dbname]" -p "[port]" -f "[name of file]"

pause
```
On mac: psql should already be added to your environment variables/path
    ```
    psql [database name] -f [path to test file]
    ```