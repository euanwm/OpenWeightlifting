# OpenWeightlifting Makefile
# Shortcuts to the most common tools should be implemented here.


# Runs the python_tools used to update the database
update_db:
	echo "Updating the database"
	@cd python_tools && pipenv run python3 backend_cli.py --update all


# Stages and commits locally all the new csv files added to the event_data folder
stage_csv:
	echo "Staging csv files"
	@git add backend/event_data/\*.csv
	@git status --p --short | grep backend/event_data
	@git commit -m "Database Update"