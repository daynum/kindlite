# Kindle highlights email service
You can upload your Kindle highlights text file and this program will parse, and save it to firestore database.

Highlights can be fetched abck via email.
And the service plans to send an email to user daily, with randomly selected highlights.

### RUN
To build the binary
```bash
go build .
```

To run the server
```bash
./kindlite
#or for windows
.\kindlite.exe
```

Access the server on `http://localhost:8000`
