# OpenWeightlifting

OK, first things first. This shares zero code or affiliation with OpenPowerlifting. However, the idea is similar. 

The overall aim is to collate as much weightliting data as possible from as many trust-worthy sources. 

Currently, those sources are national governing bodies from UK, US, and Australia. We will soon be adding competitions from the IWF.

## Local Testing

I'm not even a dev so if it doesn't work then you'll need to figure it out.
### Dockerized Website
This _**might**_ start up a live version of what is currently available.
```bash
docker-compose build
docker-compose up
```
When you get bored and want to kill it...

```bash
docker-compose down
```

### Backend-only
This did originally start as a web-scraping project which spun-off to become [sport80](https://github.com/euanwm/sport80_api) so you'll see a lot of code that bears resemblance to what would be a shoddily built API tool.
```bash
cd backend
pipenv install
pipenv run python3 backend_main.py
```

### Frontend-only
SIKE, I have no idea what i'm doing with UI stuff. It's on my lengthy list of todo's.

### License
Done this under the BSD-3-Clause license. Simply because it's what the sport80 library is under and i'm hella lazy.