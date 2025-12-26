#!/usr/bin/env bash
PATH=/bin:/sbin:/usr/bin:/usr/sbin:/usr/local/bin:/usr/local/sbin:~/bin
export PATH
# --------------------------------------------------------------
#	Project: CloudflareSpeedTest auto-update Hosts
#	Version: 1.0.4
#	Author: XIU2
#	Project: https://github.com/XIU2/CloudflareSpeedTest
# --------------------------------------------------------------

_CHECK() {
	while true
		do
		if [[ ! -e "nowip_hosts.txt" ]]; then
			echo -e "This script's purpose is to use CFST to find the fastest IP and replace Cloudflare CDN IPs in the Hosts file.\nBefore using, please read: https://github.com/XIU2/CloudflareSpeedTest/issues/42#issuecomment-768273848"
			echo -e "For first-time use, please change all Cloudflare CDN IPs in your Hosts file to a single IP."
			read -e -p "Enter that Cloudflare CDN IP and press Enter (this step is only needed once): " NOWIP
			if [[ ! -z "${NOWIP}" ]]; then
				echo ${NOWIP} > nowip_hosts.txt
				break
			else
				echo "The IP cannot be empty!"
			fi
		else
			break
		fi
	done
}

_UPDATE() {
	echo -e "Starting speed test..."
	NOWIP=$(head -1 nowip_hosts.txt)

	# You can add or modify CFST parameters here
	./cfst -o "result_hosts.txt"

	# If you want to "keep looping tests if no suitable IP is found", change both exit 0 below to _UPDATE
	[[ ! -e "result_hosts.txt" ]] && echo "CFST speed test returned 0 IPs, skipping next steps..." && exit 0

	# The following line is needed only for "keep looping tests if no suitable IP is found"
	# When a download speed limit is specified but no IP meets all conditions, CFST outputs all results
	# So when using -sl, remove the # comment symbol below and adjust the line count (e.g., if 10 IPs are tested, set value to 11)
	#[[ $(cat result_hosts.txt|wc -l) > 11 ]] && echo "CFST speed test found no IP fully meeting conditions, retesting..." && _UPDATE

	BESTIP=$(sed -n "2,1p" result_hosts.txt | awk -F, '{print $1}')
	if [[ -z "${BESTIP}" ]]; then
		echo "CFST speed test returned 0 IPs, skipping next steps..."
		exit 0
	fi
	echo ${BESTIP} > nowip_hosts.txt
	echo -e "\nOld IP: ${NOWIP}\nNew IP: ${BESTIP}\n"

	echo "Backing up Hosts file (hosts_backup)..."
	\cp -f /etc/hosts /etc/hosts_backup

	echo -e "Replacing IPs..."
	sed -i 's/'${NOWIP}'/'${BESTIP}'/g' /etc/hosts
	echo -e "Done..."
}

_CHECK
_UPDATE
