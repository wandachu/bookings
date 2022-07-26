# Bookings and Reservations

### This is a Bed-and-Breakfast booking web application that allows guests to search for available rooms and make reservations, and property owners to manage reservations. 

#### You may find and download a video demo [here](https://github.com/wandachu/bookings/tree/main/demo)

- Built in Go version 1.18
- Uses the [chi router](https://github.com/go-chi/chi/)
- Uses [alex edwards scs](https://github.com/alexedwards/scs/) for session management
- Uses [nosurf](https://github.com/justinas/nosurf)

-----
In order to build and run this application, it is necessary to
install Soda (go install github.com/gobuffalo/pop/...), create
a postgres database, fill in the correct values in database.yml,
and then run soda migrate.

To build and run the application, from the root level of the project,
execute this command:
```
go build -o bookings ./cmd/web/ && ./bookings \
-dbname=bookings \
-dbuser=postgres
```
where you have the correct entries for your database name (dbName)
and database user (dbUser)
For the full list of command flags, run ./bookings -h