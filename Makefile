include $(GOROOT)/src/Make.$(GOARCH)

.SUFFIXES: .go .$O

TARG=goals-calendar
OBJ=$(TARG).$O
GOFILES=calendar.go main.go
 
$(TARG): $(OBJ)
	$(LD) -o $@ $<

$(OBJ): $(GOFILES)
	$(GC) -o $@ $(GOFILES)

.PHONY: clean
clean:
	rm -f $(TARG)

