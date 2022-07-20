# name-ddns

Dynamic DNS for Name.com using v4 name.com API

## Usage

	docker run --rm \
		-e "NAME_DDNS_USER=namecom-user" \
		-e "NAME_DDNS_TOKEN=namecom-token" \

		-e "NAME_DDNS_CRON_0=@every 10m"
		-e "NAME_DDNS_DOMAIN_0=example.com" \
		-e "NAME_DDNS_HOST_0=subdomain" \

		-e "NAME_DDNS_CRON_1=* */5 * * * *"
		-e "NAME_DDNS_DOMAIN_1=example.com" \
		-e "NAME_DDNS_HOST_1=subdomain2" \
		-e "NAME_DDNS_TOKEN_1=other-namecom-token" \

		-e "TZ=Europe/London" \
		naxhh/name-ddns

First example will create and keep updated an A record on `subdomain.example.com.` pointing to the public IP of the network where the process is running on. This will happen every 10 minutes.

Second example runs every 5 minutes using cron format and configures an A record on `subdomain.example.com.` but uses a different API token for that particualr entry.

You can create a name.com token in https://www.name.com/account/settings/api
Supported cron formats defined at https://pkg.go.dev/github.com/robfig/cron

## Dev

Name.com api docs are in https://www.name.com/api-docs/
Cron library https://github.com/robfig/cron
