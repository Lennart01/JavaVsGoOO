JAVAC=javac
JAVA=java
SRCDIR=java_src
BINDIR=bin
SOURCES=$(wildcard $(SRCDIR)/*.java)
CLASSES=$(SOURCES:$(SRCDIR)/%.java=$(BINDIR)/%.class)
MAINCLASS=$(basename $(notdir $(SOURCES)))

all: build run

build: $(CLASSES)

run: 
	$(foreach class,$(MAINCLASS),$(JAVA) -cp $(BINDIR) $(SRCDIR).$(class);)

$(BINDIR)/%.class: $(SRCDIR)/%.java
	mkdir -p $(BINDIR)
	$(JAVAC) -d $(BINDIR) $<

clean:
	rm -rf $(BINDIR)