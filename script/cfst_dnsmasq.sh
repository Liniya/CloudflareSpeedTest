#!/usr/bin/env bash
PATH=/bin:/sbin:/usr/bin:/usr/sbin:/usr/local/bin:/usr/local/sbin:~/bin
export PATH
# --------------------------------------------------------------
#	Project: CloudflareSpeedTest auto-update dnsmasq configuration file
#	Version: 1.0.1
#	Author: XIU2,Sving1024
#	Project: https://github.com/XIU2/CloudflareSpeedTest
# --------------------------------------------------------------

_UPDATE() {
	echo -e "Starting speed test..."
	BESTIP=""
	BESTIP_IPV6="::"
	# You can add or modify CFST parameters here
	./cfst -o "result_hosts.txt"
	# To test IPv6, uncomment the following line
	#./cfst -o "result_hosts_ipv6.txt" -f ipv6.txt

	# If you want to "keep looping tests if no suitable IP is found", change both exit 0 below to _UPDATE
	[[ ! -e "result_hosts.txt" ]] && echo "CFST speed test returned 0 IPs, skipping next steps..." && exit 0

	# The following line is needed only for "keep looping tests if no suitable IP is found"
	# When a download speed limit is specified but no IP meets all conditions, CFST outputs all results
	# So when using -sl, remove the # comment symbol below and adjust the line count (e.g., if 10 IPs are tested, set value to 11)
	#[[ $(cat result_hosts.txt|wc -l) > 11 ]] && echo "CFST speed test found no IP fully meeting conditions, retesting..." && _UPDATE

	BESTIP=$(sed -n "2,1p" result_hosts.txt | awk -F, '{print $1}')
	# To test IPv6, uncomment the following line
	#BESTIP_IPV6=$(sed -n "2,1p" result_hosts_ipv6.txt | awk -F, '{print $1}')

	if [[ -z "${BESTIP}" ]]; then
		echo "CFST speed test returned 0 IPs, skipping next steps..."
		exit 0
	fi
	echo ${BESTIP} > nowip_hosts.txt
	echo -e "Best IPv4 IP: ${BESTIP}\n"
	# To test IPv6, uncomment the following line
	#echo -e "Best IPv6 IP: ${BESTIP_IPV6}\n"

    [[ -f cloudflare.conf ]] && rm cloudflare.conf

    cat site.conf | while read domain
    do
        if [[ ${domain:0:1} != "#" && ${domain} != "" ]]; then 
			echo "address=/${domain}/${BESTIP}" >> "cloudflare.conf"
			echo "address=/${domain}/${BESTIP_IPV6}" >> "cloudflare.conf"
		fi
    done

    [[ -f /etc/dnsmasq.d/cloudflare.conf ]] && rm /etc/dnsmasq.d/cloudflare.conf
    cp cloudflare.conf /etc/dnsmasq.d/cloudflare.conf
    systemctl restart dnsmasq.service
}

_UPDATE
