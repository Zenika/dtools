all: 
	cd ../src
	CGO_ENABLED=0 go build . -o ./dtools.exe

prepare:
	sudo /opt/bin/install_golang.sh 1.21.4 x86_64
	sudo apk update --no-cache

install:
	test -d $(DESTDIR)$(PREFIX)/bin || mkdir -p $(DESTDIR)$(PREFIX)/bin
	install -pm 755 $(BUILDDIR)/dtools.exe $(DESTDIR)$(PREFIX)/bin/dtools
	strip $(DESTDIR)$(PREFIX)/bin/dtools
