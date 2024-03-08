# simple-server-GA 
[![build-main](https://github.com/stianSjoli/simple-server-GA/actions/workflows/build-main.yml/badge.svg)](https://github.com/stianSjoli/simple-server-GA/actions/workflows/build-main.yml)
[![OpenSSF Scorecard](https://api.securityscorecards.dev/projects/github.com/stianSjoli/simple-server-GA/badge)](https://securityscorecards.dev/viewer/?uri=github.com/{owner}/{repo})
[![build-main](https://github.com/stianSjoli/simple-server-GA/actions/workflows/security-scans.yml/badge.svg)](https://github.com/stianSjoli/simple-server-GA/actions/workflows/security-scans.yml)


A repository with a simple application, a server, to experiment with "blue team" cyber security tools and practices. 

## SLSA
A github action workflow following SLSA, using cosign to push signed images to dockerhub.

## openSSF scorecard
openssf scan is included to analysis best practices. More about what the analysis focuses on, can be found here [OpenSSF scorecard](https://securityscorecards.dev/#the-checks).

## DAST
Owasp zap is used to scan the application, run with docker-compose. The scan image is dependent on the web application, within the compose file, and docker-compose is waiting for the exit-signal from the scan.   
