setup:
	docker-compose run web rails new . --force --database=postgresql && \
	cp -f conf/database.yml config/database.yml && \
	docker-compose build && \
	docker-compose run web rails db:create && \
	docker-compose up

