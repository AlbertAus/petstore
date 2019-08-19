## Welcome to PetStore API

This is a Golang example of https://petstore.swagger.io PetStore .

### Markdown



# Golang example of PetStore API

Go version: go version go1.12.7

1. Database version: MariaDB 10.4.6
2. Database setting in databaseSetting.go file

<!-- Need to change the following settings to your local setting -->
    serverName := "localhost:3306"  

	user := "username"  

	password := "password"  
	
	dbName := "petStore"  

3. Using app/database/petStore.sql to import sample data.
4. Only the Pet's part finished at the moment, will keep updating when I have time.