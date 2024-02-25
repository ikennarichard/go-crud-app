# CRUD APP

A crud app with Go, React and Postgres

## Notes

- List of postgres users and whether they have passwords - `SELECT usename, passwd IS NOT NULL AS has_password FROM pg_user;`
- Current postgres user - `select current_user;`
- Change user password - `ALTER USER richard WITH PASSWORD 'newpassword';`

- $ CompileDaemon -command="./MyProgram -my-options"
