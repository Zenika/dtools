# dtools

A modern-day era docker **client** with some extra features.

Empahsis here: we are talking about a cutstomized docker client. I'm nowhere close to have a docker daemon up for running.

## *dtools* vs *docker*

Basically, you use `dtools` as you use `docker`. What is more obvious is that the output is a bit changed from the official docker client, but otherwise both are similar.

Where dtools is different from docker is that some extra management commands were added here. See way below about this.

### What is *not* in `dtools`

- `network add` is not implementing the official flags, so far (will do, eventually)
- `exec` does not yet work properly.
- TODO: add more

### What is *added* in `dtools`, compared to `docker`

- default registry handling: the `dtools repo` subcommand says it all.
  Once you've used `dtools repo add` coupled with `dtools login` to that registry, using `dtools push -d` or `dtools pull -d` will automatically use that registry.
- some scripts I used to have at home to list docker images in my own registry, or tags for given image(s) are now part of the dtools client.
  See the `dtools catalog` subcommand (please note: as of now, it is not yet implemented, but is the subcommand I will implement)

## Requirements

- The foremost requirement is that you already have the docker daemon packages installed: this is needed as this software needs Docker API **v1.43** to run
- If (unsure yet, so far....) I fail to code a proper equivalent to `docker exec`, the docker client package will also be needed on your system

(the Docker API requirement is a variable in [`src/main.go`](./src/main.go#L15), this will be fixed in an ulterior version -- see [FIXME.md](./FIXME.md))

## Installing...

### Building from source

- Clone the repo
- Switch to the `src/` directory
- Run `./upgrade_pkgs.sh`
- Run `./build.sh` (have a look at the script to see offered options)

### Using the binary packages

Under the `Releases` link you should find Alpine (APK), RedHat-based (RPM) and Debian-based (DEB) packages.

### A note about building packages

The following directories are used in my own CI-CD chain at home:

- `.tito`
- `__debian`
- `__alpine`
- files `rpmbuild-deps.sh` and `dtools.spec`

Eventually, I will publish the artefacts to build the containers that use those files-directories, but as of now, they are way too customized for me to publish.
It's a bummer, those containers work oh-so-well ;)
