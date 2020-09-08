# sbt

## Project's aim
Parsing local .audit files into a well-structured data type (template for policies), which we'll use eventually to check if the policies inside the .audit file pass or not.

## Technologies used
[MySQL](https://www.mysql.com/) -> to persist the data.  
[VueJS](https://vuejs.org/) -> to read, parse the file input and communicate with the backend.  
[GO](https://golang.org/)    -> to communicate with the frontend, structure the json input into concrete data types and populate the database.  
[Wails](https://wails.app/about/) -> to compile down the project to a single executable.  

## Launching the project
First off, we need to install the dependencies: [wails](https://wails.app/gettingstarted/), [vuejs](https://vuejs.org/v2/guide/installation.html), [golang](https://golang.org/doc/install#install), [npm](https://www.npmjs.com/get-npm) [mysql](https://dev.mysql.com/downloads/installer/). We can also, optionally install a dbms, to ease our lives ([mysql-workbench](https://www.mysql.com/products/workbench/))  
After installing everything, `cd` into the root of the project and run  
`wails serve`  
`cd frontend`  
`npm run serve`  
`cd ../build && wails build -p` -> To pack the project into a single binary  
`./app_name` -> To run the binary  
