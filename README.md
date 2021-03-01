prerequisites
------------
install docker

https://docs.docker.com/engine/install/

build go app
-----------
clone the repo
cd go_proj
docker build -t goapp .
docker run -d -p 8080:8080 --name mygoapp goapap.

if 8080 port is aalready lissten in your system change your host port like below.

docker run -d -p 8081:8080 --name mygoapp goapap 

this app contaians two end points aaaccess the using below endpoints

http://<hostport>:8080/ --> welcomepaage
http://<hostport>:8080/api/encrypt --> it will print encrypt json
http://<hostport>:8080/api/decrypt --> print decrypt json

samaple json already pass in the code:
{ID: "1", Name: "IBM"} 

