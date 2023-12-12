# Preparing Repo

## Gin over Mux

- [Gin](https://gin-gonic.com/)
- Nesting Routes in Mux is tough
- Applying Middlewares to specific routes is not so simple in Mux

## DB Model

- Table
  - User
    - Id
    - Name
  - Plane
    - Id
    - Name
    - Source
    - Destination
  - Seats
    - PlaneId
    - SeatNo
  - Bookings
    - Id
    - UserId
    - PlaneId
    - SeatNo
    - BookingDateTime

## Routes

- [app/routers/main.go](app/routers//main.go) -> Declare a master router
- App/Routers/[route-Group].routes.go -> Declare a router group and individual routes
  - E.g. User Routes are present in [app/routers/user.routes.go](app/routers/user.routes.go)
