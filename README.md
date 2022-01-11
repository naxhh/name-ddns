# name-ddns

Dynamic DNS for Name.com using v4 name.com API

## Usage

	docker run --rm \
		-e "NAME_DDNS_USER=namecom-user" \
		-e "NAME_DDNS_TOKEN=namecom-token" \
		-e "NAME_DDNS_DOMAIN=example.com" \
		-e "NAME_DDNS_HOST=subdomain" \
		-e "NAME_DDNS_UPDATE_EVERY_MINUTES=10" \
		-e "TZ=Europe/London" \
		naxhh/name-ddns

This example will create and keep updated an A record on `subdomain.example.com.` pointing to the public IP of the network where the process is running on.

You can create a name.com token in https://www.name.com/account/settings/api

## Dev

Name.com api docs are in https://www.name.com/api-docs/
