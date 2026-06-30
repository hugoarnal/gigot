GO	?=	go

BINARY	:=	gigot

$(BINARY):
	$(GO) build -o $(BINARY) gigot.go

run:
	$(GO) run gigot.go

fclean:
	$(RM) $(BINARY)

re: fclean $(BINARY)

.PHONY:	run fclean re
