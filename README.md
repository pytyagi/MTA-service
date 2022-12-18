# MTA-HOST_OPTIMIZER

## Service to track the inefficent servers

* Steps to run the service
```
$ git clone https://github.com/pytyagi/MTA-service.git
$ cd MTA-service
$ make run
```
* Update .env file for X (threshold limit) and mock service URL.
* Use http://localhost:8080/hosts with HTTP GET to get the response.

* To get Coverage run the below command and run index.html in browser.
```
make local-cover
```
* To mock the service for IP Data, I have used data in test.json file in repo.
* CI/CD is configured using Github Actions on each git push pipeline for test and build will run.