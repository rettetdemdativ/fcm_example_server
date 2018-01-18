# Example web server using Firebase Cloud Messaging

This server was originally created for a talk at the Vienna Go Meetup. Additionally, it serves as an example implementation for anyone trying to create a web server in Go that allows sending data messages to client applications via Firebase Cloud Messaging.

The implementation currently uses the [legacy HTTP API](https://firebase.google.com/docs/cloud-messaging/http-server-ref) and a permanent API key which can be acquired through the Firebase Console. 

## Acquiring the legacy API key
1. Go to the [Firebase Console](https://console.firebase.google.com).
2. Select your project (set one up if you haven't yet)
3. Go to the project's settings and select the *Cloud Messaging* tab. There you'll find your API key.