 # Skills presentation test
 
 ## Requirements
 
 - docker
 
 ## Dependencies
 
 - [gorilla/mux](github.com/gorilla/mux) it provides an easy to configure routing system compatible whit net/http
 
 thanks to go module there is no need to go get -t each package
    
 ``` go mod download```
    
 it will download all the modules if not already available.
 
 ## Usage
 
 clone this repo and enter the repo and run the following commands,
 depending on your system privileges on the host machine it may require 
 to insert your credentials if you have access to sudo privileges or
 be restricted.  
 
 The second command will assign 4 cpu to the docker container and 50MB of ram
 
 ```
    sudo docker build -t apptest:1.0 .
    sudo docker run -d -p 8001:8123 -it --cpus="4" --memory=500m apptest:1.0
 ```

At this stage the app will be available at 

```
    http://127.0.0.1:8001/exchange/{TICKER}
```

substitute "TICKER" with the one you wish the response USD or EUR are the only currently supported.


## Conclusions and future implementations

I have tried to be as much as concise and leans as possible in this skill set showcase.
As future development I would like to list the following:

 - a middleware that handle gracefully any potential panic that may occur during the execution
  of any of the api calls provided by the app to prevent a single panic from halting the whole application.
 
 - a middleware that provide caching of the endpoint in case a tolerance of at least 100ms its allowed on the refresh 
   of the data.
 
 - a middleware that limit the access to the api to prevent a denial of service attack and/or over consumption of resources. 
 
 - expand on logging system in case of error to keep track of any warning or incident happened during the execution.


# TODO Improve

the function GetRateByDate and the function was too big
- new features, integrate with a token jwt received from auth0
- store user currency researches (user, name, list of currency conversions)
- move to buffalo with graphql plugin for apollo

