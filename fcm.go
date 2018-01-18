package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const fcmServerURL = "https://fcm.googleapis.com/fcm/send"
const keyFilePath = "./api.key"

var fcmKey string

// Prepares the Firebase Cloud Messaging connection by setting the API key to
// use during authorization.
// A key can be claimed at
// https://console.firebase.google.com/project/<project-id>/settings/cloudmessaging/
func prepareFCM() {
	data, err := ioutil.ReadFile(keyFilePath)
	if err != nil {
		panic(err)
	}
	fcmKey = string(data)
}

// Global map of Firebase Cloud Messaging tokens to keep track of the ones
// already registered with the server. Any message sent with sendFCMMessageToAll
// will be sent to all tokens in this map.
var fcmTokenMap = map[string]bool{}

// Default Firebase Cloud Messaging message body. Contains the required fields
// for sending a data message, multi-cast and uni-cast.
// Reference: https://firebase.google.com/docs/cloud-messaging/http-server-ref
type fcmMessage struct {
	To              string      `json:"to,omitempty"`
	RegistrationIDs []string    `json:"registration_ids,omitempty"`
	Data            interface{} `json:"data,omitempty"`
}

// Sends the given message in a pre-defined JSON structure to all clients that
// registered their tokens.
func sendFCMMessageToAll(message string) error {
	var ids []string
	for k := range fcmTokenMap {
		ids = append(ids, k)
	}

	m := fcmMessage{
		RegistrationIDs: ids,
		Data:            map[string]string{"message": message},
	}

	jd, err := json.Marshal(&m)
	if err != nil {
		log.Printf("Failed to marshal JSON: %s", err.Error())
		return err
	}

	log.Printf("FCM Message: %s", string(jd))

	r, err := http.NewRequest("POST", fcmServerURL, bytes.NewReader(jd))
	if err != nil {
		log.Printf("Failed to create new HTTP request: %s", err.Error())
	}

	// Add the Authorization header with the API key
	r.Header.Set("Authorization", fmt.Sprintf("key=%s", fcmKey))

	r.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		log.Printf("Request to %s failed: %s", fcmServerURL, err.Error())
		return err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("Request returned status %s:\n\tHeader: %s\n\tBody: %s",
		resp.Status,
		resp.Header,
		body,
	)

	return nil
}
