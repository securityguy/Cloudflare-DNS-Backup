# Cloudflare-DNS-Backup

This utility accepts a directory and Cloudflare token on the command line:

    Cloudflare-DNS-Backup <directory> <token>

Using this information, it queries the Cloudflare API, obtains a list of
zones, and exports DNS records for each zone to a separate file in the specified directory.

Cloudflare "Export DNS Records" is used because records to facilitate
re-importing them using the API.

Importation has not yet been tested.

Please see https://api.cloudflare.com/#dns-records-for-a-zone-export-dns-records for details.

This application is a work in progress -- use at your own risk.

Please see the LICENCE file for licence information.

To compile, clone this repo, enter the directory, and run:

    go build
