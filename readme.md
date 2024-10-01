#### **Setup**

##### Docker setup

- change docker-compose.yml POSTGRES_USER and POSTGRES_PASSWORD to the required user and password
- add  ***.env*** with the right DB credentials and hostname to the root folder of the project
- run docker-compose build and docker-compose up
- connect to the docker postgerSQL using your SQL client or if you dont have one download the community version of [dbeaver](https://dbeaver.io/ "dbeaver")
- then copy the contents of ***sqlTableArch.sql*** and paste it at your SQL script editor and run it.

##### No docker setup

- install go1.22.5
- install postgresql 
- add  ***.env*** with the right DB credentials and hostname to the root folder of the project
- create a database using postgresql name it Contact
- select Contact as your Database
- then copy the contents of ***sqlTableArch.sql*** and paste it at your SQL script editor
- then *run go .*  to the root folder of the project

------------

#### **How does it work**

##### Docker
- run postman, insomnia or any API testing app you have
- to add contacts add this url [localhost:5000/addContact](http://localhost:5000/addContact "localhost:5000/addContact") in your API testing app with a method **POST** and with a data:
	> 	{
		"full_name": "Alex Bell", 
		"email": "alex@bell-labs.com",
		"phone_numbers":["03 8578 6688", "1800728069"]
		}
	OR run this command in your terminal(windows)
	>
		curl --location localhost:5000/addContact --header "Content-Type: application/json" --data-raw "{\"full_name\": \"Alex Bell\", \"email\": \"alex@bell-labs.com\",\"phone_numbers\":[\"03 8578 6688\", \"1800728069\"]}"
	>
- to see if you added something go your browser and go to this link [localhost:5000/getAll](http://localhost:5000/getAll "localhost:5000/getAll")

##### No Docker
- just change the localhost:5000 to localhost:8080
