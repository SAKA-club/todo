REM Run psql
"C:\Program Files\PostgreSQL\14\bin\psql.exe" -h "localhost" -U "postgres" -d "todoDB" -p "5432" -f "testdata.sql"

pause