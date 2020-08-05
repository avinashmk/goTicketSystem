# goTicketSystem [Under dev]
 
A robust application to reserve/book Tickets.

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
   |   |---repo_structure.png: Repo structure diagram.  
   |  
   |---internal  
   |   |  
   |   |---consts  
   |       |  
   |       |---consts.go: Holds all hard-coded values in one place.  
   |  
   |---core  
   |   |  
   |   |---core.go: Starts core processes -- store(for db),  
   |                                         handler(for server connections),  
   |                                         housekeeping(for maintenance).  
   |  
   |---handler  
   |   |  
   |   |---handlercontrol.go  
   |   |---common.go:  
   |   |---signin.go:  
   |   |---signup.go:  
   |   |---searchtrain.go:  
   |   |---makereservation.go:  
   |   |---viewreservation.go:  
   |   |---canclereservation.go:  
   |  
   |---housekeeping  
   |   |  
   |   |---housekeeping.go:  
   |  
   |---model  
   |   |  
   |   |---general.go  
   |   |---menu.go  
   |  
   |---store  
   |   |  
   |   |---store.go  
   |   |---userdoc.go  
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
   |  
   |---util  
   |   |  
   |   |---util.go: provides common utility methods that may/may not be specific to program  
   |  
   |---logs  
```  
  