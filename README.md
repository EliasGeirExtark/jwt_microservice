# JWT Auth Microservice
The scope of this service is a module to include with docker inside your projects, to use jwt authentication.

## Install
1) Copy and rename the file .env.dist in .env file. Then open it and compile with your data. Be sure that the db data is correct and the database can communicate with your program. This release did not include the database setting, because We assume that this module will be added to another project;

## Endpoints
In this section you can see all the endpoints of this microservice, with a small descripiton, a example and the results
### Login
This function is used to make the login. When you insert the correct credentials, the result will be an access token, a refresh token and the data.
The data represents only the data content inside the jwt token, or the key of the user.
##### Endpoint
```
POST {host}/api/login
```
##### Body
```
{
    "user": "user_credential",
    "password": "your_password"
}
```
##### Response
```
{
    "data": "user_id"
    "access_token": "<your_access_token>"
    "refresh_token": "<your_refresh_token>"
}
```
### Refresh
Without make the login after every expire time, you can use the refresh token to make a secured login without any credential. You need to send the access and refresh token to refresh this two with 2 newer tokens.
##### Endpoint
```
POST {host}/api/refresh
```
##### Body
```
{
    "access_token": "<your_access_token>"
    "refresh_token": "<your_refresh_token>"
}
```
##### Response
```
{
    "data": "user_id"
    "access_token": "<your_access_token>"
    "refresh_token": "<your_refresh_token>"
}
```
## Dependences
- [gorilla/mux](https://github.com/gorilla/mux)
- [validator](https://github.com/go-playground/validator)