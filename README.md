# httpController
go based linux daemon controlled via JSON

Usage:
Can be started just via:
./httpController - will listen on port 8080 on the default IP Address/NIC (or local loopback 127.0.0.1)

Commands to start applications can be sent via JSON in the following format:

To Start Processes:
{"command":"ping","arg1":"127.0.0.01","cwd":"/home/user/","state":"started"}

Example via CuRL

curl -H "Content-Type: application/json" -d '{"command":"ping","arg1":"127.0.0.01","cwd":"127.0.0.1","state":"started"}' http://localhost:8080/

To Stop Processes:

{"command":"ping","arg1":"127.0.0.01","cwd":"/home/user/","state":"stopped"} 

Example via CuRL

curl -H "Content-Type: application/json" -d '{"command":"ping","arg1":"127.0.0.01","cwd":"127.0.0.1","state":"stopped"}' http://localhost:8080/

Killing a process started via the httpController (ie kill -9) will cause the httpController to startup the process - processes can only be stopped via the state:stopped command as above
