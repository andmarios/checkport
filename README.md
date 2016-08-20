# checkport #

A very simple app to test if a (TCP) port is available to bind to.

Basic example for binding to port 8080 in all interfaces:

    checkport -port 8080

Test if we can bind to port 8000 in localhost:

    checkport -address localhost -port 8000
