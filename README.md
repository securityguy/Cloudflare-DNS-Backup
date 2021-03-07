# Cloudflare-DNS-Backup

This utility accepts a directory and Cloudflare token on the command line:

    Cloudflare-DNS-Backup <directory> <token>

Using this information, it queries the Cloudflare API, obtains a list of
zones, and exports DNS records for each zone to a separate file in the specified directory.

Cloudflare's "Export DNS Records" method is used to facilitate re-importing the data using the API if required. Importation has not yet been tested.

Please see https://api.cloudflare.com/#dns-records-for-a-zone-export-dns-records for details.

This application is a work in progress -- use at your own risk.

Please see the LICENCE file for licence information.

To compile, clone this repo, enter the directory, and run:

    go build
