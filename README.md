# goTicketSystem [Under dev]
 
A robust and autonomous application to manage Tickets.

Features:
- Handle concurrent users.
- Encrypted passwords.
- MongoDB.

Structure:  
```  
   Repo  
   |  
   |---command  
   |   |  
   |   |---TicketSystem  
   |       |  
   |       |---main.go: Produces executable  
   |  
   |---docs  
   |   |  
   |   |---db_design.png: Database design diagram.
   |  
   |---internal  
   |   |  
   |   |---consts  
   |   |   |  
   |   |   |---consts.go: Holds all hard-coded values in one place.  
   |   |  
   |   |---core  
   |   |   |  
   |   |   |---core.go: Starts core processes -- store(for db),  
   |   |                                         servercontrol(for server connections),  
   |   |                                         housekeeping(for maintenance).  
   |   |  
   |   |---housekeeping  
   |   |   |  
   |   |   |---housekeeping.go:  
   |   |   |---charts.go:  
   |   |   |---tickets.go:  
   |   |  
   |   |---model  
   |   |   |  
   |   |   |---general.go:  
   |   |   |---menu.go:  
   |   |  
   |   |---server  
   |   |   |  
   |   |   |---server.go: Starts up web server and sets up all handlers  
   |   |   |  
   |   |   |---session  
   |   |   |   |  
   |   |   |   |---session.go:  
   |   |   |  
   |   |   |---handler  
   |   |       |  
   |   |       |---common.go:  
   |   |       |---signin.go:  
   |   |       |---signup.go:  
   |   |       |---signoff.go:  
   |   |       |---searchtrain.go:  
   |   |       |---makereservation.go:  
   |   |       |---viewreservation.go:  
   |   |       |---cancelreservation.go:  
   |   |       |---addtrainschema.go:  
   |   |       |---removetrainschema.go:  
   |   |       |---updatetrainschema.go:  
   |   |       |---viewtrainschema.go:  
   |   |  
   |   |---store  
   |       |  
   |       |---store.go:  
   |       |---userdoc.go:  
   |       |---schemaDoc.go:  
   |       |---chartDoc.go:  
   |  
   |---logger  
   |   |  
   |   |---logger.go: Provides logging utility  
   |  
   |---web  
   |   |  
   |   |---static  
   |   |   |  
   |   |   |---index.html  
   |   |  
   |   |---templates  
   |       |  
   |       |---menu.html  
   |       |---signin.html  
   |       |---signup.html  
   |       |---addtrainschemaform.html  
   |  
   |---util  
   |   |  
   |   |---util.go: provides common utility methods that may/may not be specific to program  
   |  
   |---logs  
```  
  