# OpenWeightlifting Makefile
# Shortcuts to the most common tools should be implemented here.

# Builds the backend server executable
.PHONY: build_backend
build_backend:
	cp -r event_data/ backend/event_data/
	cd backend && go build -o backend

# Builds the frontend files
.PHONY: build_frontend
build_frontend:
	@cd scripts && pipenv install && pipenv run python3 fe_builder.py
	cd frontend && npm install

.PHONY: check_db
DB ?= ""
check_db:
	@cd python_tools && pipenv run python3 check_db.py $(DB)

.PHONY: generate-docs
generate-docs:
	echo "Generating docs..."
	cd backend && swag init --parseDependency --parseInternal

# Removes build files.
.PHONY: clean
clean:
	rm -f backend/backend
	rm -rf backend/event_data

# Removes build files plus cached dependencies.
.PHONY: veryclean
veryclean: clean
	rm -rf frontend/node_modules
