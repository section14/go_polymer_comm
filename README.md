# go-polymer-comm

CMS based off of Go (Golang) for the backend api and Google Polymer. Database for this project is Google Datastore (via Google App Engine). The backend api is designed to be consumed by any front end, and doesn't have any dependencies on Polymer to run.

Backend Go packages are located here: https://github.com/section14/go_polymer_comm_pkg

Use this terminal command to launch instead of `goapp serve`: 

--datastore_path=~/go_apps/data/datastore.db dispatch.yaml backend/app.yaml frontend/app.yaml
