# goTicketSystem [Under dev]
 
A robust application to reserve/book Tickets.

Features:
- Handle concurrent users.
- Encrypted passwords.
- MongoDB.

Structure:
 - command/TicketSystem
    main.go: Produces executable

 - docs
    repo_structure.png: Repo structure diagram.

 - internal
    contains all program internal packages.

 - internal/consts
    consts.go: Holds all hard-coded values in one place.

 - internal/core
    core.go: Starts core processes -- store(for db), handler(for server connections), housekeeping(for maintenance).

 - internal/handler
    handlercontrol.go: 
    common.go:
    signin.go:
    signup.go:
    searchtrain.go:
    makereservation.go:
    viewreservation.go:
    canclereservation.go

 - internal/housekeeping
    housekeeping.go:

 - internal/model
    general.go:
    menu.go:

 - internal/store
    store.go:
    userdoc.go:

 - logger
    logger.go: provides logging methods.

 - logs

 - util
    util.go: provides common utility methods that may/may not be specific to program.

 - web
    contains web contenting related data

 - web/static
    index.html:

 - web/templates
    signin.html:
    signup.html:
    menu.html: