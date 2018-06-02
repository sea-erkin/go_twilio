# go_twilio
Super simple twilio interface which allows you to send an SMS message


```
go get github.com/sea-erkin/go_twilio
```

Initiate client
```
twilioClient, err := twilio.New("Your account id", "Your account password", verboseBool)
```

Send SMS message
```
err = TwilioClient.SendMessage("11234567890", "11234567890", "Test message body")
```
