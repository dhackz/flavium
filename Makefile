#!make
include .env
export $(shell sed 's/=.*//' .env)

build: flavium-dashboard flavium-backend
	  $(MAKE) -C flavium-dashboard/
	  $(MAKE) -C flavium-backend/
	  
dashboard: flavium-dashboard 
	$(MAKE) -C flavium-dashboard/
	
backend: flavium-backend
	$(MAKE) -C flavium-backend/
