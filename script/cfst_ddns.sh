#!/usr/bin/env bash
PATH=/bin:/sbin:/usr/bin:/usr/sbin:/usr/local/bin:/usr/local/sbin:~/bin
export PATH
# --------------------------------------------------------------
#	Project: CloudflareSpeedTest auto-update DNS record
#	Version: 1.0.5
#	Author: XIU2
#	Project: https://github.com/XIU2/CloudflareSpeedTest
# --------------------------------------------------------------

_READ() {
	[[ ! -e "cfst_ddns.conf" ]] && echo -e "[Error] Configuration file does not exist [cfst_ddns.conf] !" && exit 1
	CONFIG=$(cat "cfst_ddns.conf")
	FOLDER=$(echo "${CONFIG}"|grep 'FOLDER='|awk -F '=' '{print $NF}')
	[[ -z "${FOLDER}" ]] && echo -e "[Error] Missing configuration item [FOLDER] !" && exit 1
	ZONE_ID=$(echo "${CONFIG}"|grep 'ZONE_ID='|awk -F '=' '{print $NF}')
	[[ -z "${ZONE_ID}" ]] && echo -e "[Error] Missing configuration item [ZONE_ID] !" && exit 1
	DNS_RECORDS_ID=$(echo "${CONFIG}"|grep 'DNS_RECORDS_ID='|awk -F '=' '{print $NF}')
	[[ -z "${DNS_RECORDS_ID}" ]] && echo -e "[Error] Missing configuration item [DNS_RECORDS_ID] !" && exit 1
	KEY=$(echo "${CONFIG}"|grep 'KEY='|awk -F '=' '{print $NF}')
	[[ -z "${KEY}" ]] && echo -e "[Error] Missing configuration item [KEY] !" && exit 1
	EMAIL=$(echo "${CONFIG}"|grep 'EMAIL='|awk -F '=' '{print $NF}')
	[[ -z "${EMAIL}" ]] && echo -e "[Info] Missing configuration item [EMAIL], switching from [API Key] to [API Token] mode!"
	TYPE=$(echo "${CONFIG}"|grep 'TYPE='|awk -F '=' '{print $NF}')
	[[ -z "${TYPE}" ]] && echo -e "[Error] Missing configuration item [TYPE] !" && exit 1
	NAME=$(echo "${CONFIG}"|grep 'NAME='|awk -F '=' '{print $NF}')
	[[ -z "${NAME}" ]] && echo -e "[Error] Missing configuration item [NAME] !" && exit 1
	TTL=$(echo "${CONFIG}"|grep 'TTL='|awk -F '=' '{print $NF}')
	[[ -z "${TTL}" ]] && echo -e "[Error] Missing configuration item [TTL] !" && exit 1
	PROXIED=$(echo "${CONFIG}"|grep 'PROXIED='|awk -F '=' '{print $NF}')
	[[ -z "${PROXIED}" ]] && echo -e "[Error] Missing configuration item [PROXIED] !" && exit 1
}

_UPDATE() {
	# You can add or modify CFST parameters here
	./cfst -o "result_ddns.txt"

	# Check if result file exists; if not, results are 0
	[[ ! -e "result_ddns.txt" ]] && echo "CFST speed test returned 0 IPs, skipping next steps..." && exit 0

	CONTENT=$(sed -n "2,1p" result_ddns.txt | awk -F, '{print $1}')
	if [[ -z "${CONTENT}" ]]; then
		echo "CFST speed test returned 0 IPs, skipping next steps..."
		exit 0
	fi
	# If EMAIL variable is empty, use API token mode
	if [[ -n "${EMAIL}" ]]; then
		# API Key method (global permissions)
		curl -X PUT "https://api.cloudflare.com/client/v4/zones/${ZONE_ID}/dns_records/${DNS_RECORDS_ID}" \
			-H "X-Auth-Email: ${EMAIL}" \
			-H "X-Auth-Key: ${KEY}" \
			-H "Content-Type: application/json" \
			--data "{\"type\":\"${TYPE}\",\"name\":\"${NAME}\",\"content\":\"${CONTENT}\",\"ttl\":${TTL},\"proxied\":${PROXIED}}"
	else
		# API Token method (custom permissions)
		curl -X PUT "https://api.cloudflare.com/client/v4/zones/${ZONE_ID}/dns_records/${DNS_RECORDS_ID}" \
			-H "Authorization: Bearer ${KEY}" \
			-H "Content-Type: application/json" \
			--data "{\"type\":\"${TYPE}\",\"name\":\"${NAME}\",\"content\":\"${CONTENT}\",\"ttl\":${TTL},\"proxied\":${PROXIED}}"
	fi
}

_READ
cd "${FOLDER}"
_UPDATE
