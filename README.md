# Pet Adoption Tracker

This is a simple GUI app in Go built using the Fyne GUI framework.

The app allows a user to add animals ready for adoption into a database, as well
as pick an animal of desired type (e.g. dog) to adopt from the database (also
removing it from the database).

## Getting Started: Building and Running the App

### Requirements:

- Go (version >= 1.19 for Fyne)
- C compiler (e.g. `gcc`) for Fyne
- sqlite3

### Cloning Repo and Building:

```bash
git clone https://github.com/Kevonosdiaz/pet-adoption-tracker
cd pet-adoption-tracker
go build .
./pet-adoption-tracker
```

The `CGO_ENABLED=1` environment variable may have to be passed if building
fails, as well as ensuring `gcc` and sqlite are available on the system.

### Running Tests:

```
go test
```
