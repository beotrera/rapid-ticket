# Rapid Ticket 🎫

This project is a RESTful API that allows booking tickets for events, managing seat availability and displaying details of shows.

## Requirements

- Go 1.18 or higher
- SQLite database
- Dependencies managed with `go mod`

### Clone the repository

Clone this repository to your local machine:

```bash
git clone https://github.com/tu-usuario/tu-repo.git
cd tu-repo
```

## Install dependencies

This project uses Go Modules for dependency management. To install all necessary dependencies, run:

```bash
go mod tidy
```

## Run

To run the server, use the following command

```bash
go run main.go
```

The server will run on port **:3000** by default.

## Users

Para usar los endpoints con basic auth se puede usar el usuario:

- **Username**: test@test.com
- **Password**: Password123!

## Endpoints

# Get - Get Shows

```bash
curl --location 'http://localhost:3000/shows' \
--header 'x-api-key: ro4U9iMcJm0FxjMd4uS80ip5L1XqyOWLDuQCQ6jT4BA=' \
--header 'x-api-secret: 5b3c2299886c3d8bfad03735eac6665a1035432f9752923a81366562eada38b4dfb6d4b930c9f5e56f7cce4acdb5e88a82a53293219de26f284c5103b5f49e23.f9951364ca631f64' \
--header 'Content-Type: application/json' \
--header 'Cookie: sessionID=s%3A3wC02_z_7F4xpY2nPKz4kYh5x9kNGc9Z.v4gD6fwXhyMccKQExSrXBgTj3mXrbNEPJ4oGGF5h%2FOk; sessionID=s%3AiAz8mSTPBUhdhaCxOFtgynOw66fkqIBd.vIjTp6FGMpP9QYyY6c1%2FrBMcb6%2BUY3hX8dl%2FnZ3Un8o' \
--data ''
```

# Post - Create Reservation

```bash
curl --location 'http://localhost:3000/reservations' \
--header 'Content-Type: application/json' \
--header 'Authorization: Basic dGVzdEB0ZXN0LmNvbTpQYXNzd29yZDEyMyE=' \
--header 'Cookie: sessionID=s%3AiAz8mSTPBUhdhaCxOFtgynOw66fkqIBd.vIjTp6FGMpP9QYyY6c1%2FrBMcb6%2BUY3hX8dl%2FnZ3Un8o' \
--data '{
"showId":3,
"dni":"38827601",
"name":"test",
"seats": [
{
"sectionId":5,
"seat":"A1"
}
]
}'
```
