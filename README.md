# App Props
java like application properties for Golang. Hassle free config management for deployment to the cloud.

## Overview 

Engineers coming from the Java spring boot background tend to put a lot of env config into ```application.properties``` and are used to it. this is a 
micro-library to help you out with managing the configs in the same old ```application.properties```. 
Managing multiple profiles is very easy just add profile name by application ```application-xxx.properties```
by default application properties will be loaded which can then be guided by the profile to be used. 
All profiles will by default pick variables from ```application.properties``` if not found in ```application-xxx.properties```

To set environment variables use 
```SOME_CONFIG_KEY=${ENV_KEY_NAME:ENV_DEFAULT_VALUE_IF_NOT_FOUND}```

### How to Use 

```go 

func main() {
	config := appProps.UseProps("resources/application.properties") // your config chain is loaded here 
        // pass it with your function wherever you want to use it. 
        // to get any value from the config use 
        conigValue := config.Get("KEY_NAME") // returns a string 
}
```
