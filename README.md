# OpenWeightlifting

[![codecov](https://codecov.io/gh/euanwm/OpenWeightlifting/branch/development/graph/badge.svg?token=CX7H10ZNLM)](https://codecov.io/gh/euanwm/OpenWeightlifting)

## Local Testing
We've added a docker-compose file to make it easier to test locally. This will spin up a local instance of the backend and frontend services. In production, these services are deployed separately.
```bash
docker-compose build
docker-compose up
```
When you get bored and want to kill it...

```bash
docker-compose down
```

### Backend-only
When launching the backend service you'll need to toggle the CORS flag which is done be adding the 'local' argument when calling the executable.
```bash
go build backend.go
./backend local
```

### Frontend-only
We prefer to use npm for the frontend stuff.
```bash
npm install
npm run dev
```

### Updating the database
To pull the latest results from the all relevant federations, you'll need to run the following command from the python_tools directory:
```bash
pipenv install
pipenv run python backend_cli.py --update all
```

### License
Done this under the BSD-3-Clause license. Simply because it's what the sport80 library is under and i'm hella lazy.