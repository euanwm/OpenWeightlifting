# OpenWeightlifting
This is the monorepo for the OpenWeightlifting.org project. The aim of this project is to build a database of the latest Olympic Weightlifting results from all around the world. This originally started from a scraping tool and quickly grew into what you see here now. All the results within the database were pulled directly from the event results pages from the National Governing Body of that nation. We try to avoid manual data entry so this is all done with our tooling written in Python.

There's always some questions around certain design decisions in this project so we'll address them now.

### Why Golang for the backend? 
Originally it was Python but the build time was terrible and the response times were slow. Not only that but the memory usage was high. Golang was chosen because it's fast, has a low memory footprint and the build times are quick. It's also a language that's easy to pick up and learn. We migrated from Python to Golang within in a week of picking up the language.

### Why NextJS for the frontend?
This was a bit of a no brainer. We wanted to use React and NextJS (TS) is a great framework for it. The amount of features around rounting, server side rendering and static site generation is great. While it can also serve as a backend, we've chosen to keep the backend and frontend separate due the performance benefits of having a dedicated backend.

## Local Testing
Majority of the contributors on this are FE developers, so we've containerised the front and backend portions of the project. To get going quickly with the project, you'll need to have docker installed.

### Frontend Development (NextJS)
From the root of the project, run the following commands to spin up a backend container and launch the frontend.
```bash
docker compose up -d backend
cd frontend
npm install
npm run dev
```
While the backend service is running, you'll also be able to run the FE API call tests against it.
```bash
npm run test
```

Once you're done, you can stop the backend container with the following command.
```bash
docker compose down
```

### Backend Development (Golang)
When launching the backend service you'll need to toggle the CORS flag which is done be adding the 'local' argument when calling the executable.
```bash
go build backend.go
./backend local
```

## Database Management (Python)
To pull the latest results from the all relevant federations, we've added a Makefile with a few commands to make it easier. You'll need to have pipenv installed to run the commands.
### Pulling the latest results
```bash
make update_database
```

### Staging the latest results
This came about so if you have a messy amount of unstaged changes, you can stage them all and then commit them in one go.
```bash
make stage_csv
```

### Checking the database
In order to reduce the amount of checks at runtime, we've added a check to make sure the database is in a good state. This will check for duplicate entries and missing data.
```bash
make check_db
```

# License
Done this under the BSD-3-Clause license. Simply because it's what the sport80 library is under and i'm hella lazy.

# Contributing
If you want to contribute, please do! We're always looking for more people to help out. If you're not sure where to start, check out the issues page and join the discord server.
https://discord.gg/kqnBqdktgr