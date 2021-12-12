# name-ddns

Dynamic DNS for Name.com using v4 name.com API

## Usage

		export NAME_DDNS_USER="name.com user"
		export NAME_DDNS_TOKEN="name.com api token"
		export NAME_DDNS_UPDATE_EVERY_MINUTES="10"
		export NAME_DDNS_DOMAIN="tolstoy.eu"
		export NAME_DDNS_HOST="example" # fqdn will be example.tolstoy.eu.
		
		name-ddns

You can create a token in https://www.name.com/account/settings/api


## Dev

Name.com api docs are in https://www.name.com/api-docs/
