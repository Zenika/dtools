<H1>dtools</H1>
___

A modern-day era docker **client** with some extra features.<br>
Empahsis here: we are talking about a cutstomized docker client. I'm nowhere close to have a docker daemon up for running.<br>


<H2>Conventional docker client vs dtools</H2>
Please see MAPPINGS.md to see which docker command are implemented so far, and how it translates with dtools, and, the roadmap.
<H2>How to use</H2>
Basically, you use `dtools` as you use `docker` . What is more obvious is that the output is a bit changed from the official docker client, but otherwise both are similar.

Where dtools is different is that some extra management commands were added here. See way below about this.

<H2>Requirements</H2>
- The foremost requirement is that you already have the docker daemon packages installed : this is needed as this software needs Docker API **v1.43** to run
- If (unsure yet, so far....) I fail to code a proper equivalent to `docker exec`, the docker client package will also be needed on your system

(the Docker API requirement is a variable in src/main.go, this will be fixed in an ulterior version-- see FIXME.md)
<H2>Installing....</H2>
___
<H3>Building from source</H3>
- Clone the repo<br>
- Switch to the `src/` directory<br>
- Run `./upgrade_pkgs.sh`<br>
- Run `./build.sh` (have a look at the script to see offered options)<br>

<H3>Using the binary packages</H3>
Under the `Releases` link you should find Alpine (APK), RedHat-based (RPM) and Debian-based (DEB) packages

<H3>A note about building packages</H3>
The following directories are used in my own CI-CD chain at home:<br>
- .tito<br>
- __debian<br>
- __alpine<br>
- files rpmbuild-deps.sh and dtools.spec<br>

Eventually, I will publish the artefacts to build the containers that use those files-directories, but as of now, they are way too customized for me to publish.
It's a bummer, those containers work oh-so-well ;)

<H2>*dtools* vs *docker*</H2>
___
<H3>What is **not** vs **dtools**</H3>
- `network add` is not implementing the official flags, so far (will do, eventually)
- `exec` actually forks a shell to `docker exec`
- TODO: add more

<H3>What is *added* in `dtools`, compared to `docker`</H3>
- default registry handling: the `dtools repo` subcommand says it all.<br>
    Once you've used `dtools registry add` coupled with `dtools login` to that registry, using `dtools push -d` or `dtools pull -d` will automatically use that registry.<br>
- some scripts I used to have at home to list docker images in my own registry, or tags for given image(s) are now part of the dtools client<br><br>
See the `dtools catalog` subcommand (please note: as of now, it is not yet implemented, but is the subcommand I will implement)