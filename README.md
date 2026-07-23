# Pet Adoption Tracker

This is a simple GUI app in Go built using the Fyne GUI framework.

## Getting Started: Building and Running the App

### Requirements:

- Go (version >= 1.19 for Fyne)
- C compiler (e.g. gcc) for Fyne and `mattn/go-sqlite3` (a CGO package)
- sqlite3

### Cloning Repo and Building:

```bash
git clone https://github.com/Kevonosdiaz/pet-adoption-tracker
cd pet-adoption-tracker
# Note: CGO_ENABLED=1 environment variable may need to be explicitly passed for `go build .`
go build .
./pet-adoption-tracker
```
