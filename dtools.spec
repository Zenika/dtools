%define debug_package   %{nil}
%define _build_id_links none
%define _name dtools
%define _prefix /opt
%define _version 00.73.00
%define _rel 0
%define _arch x86_64
%define _binaryname dtools

Name:       dtools
Version:    %{_version}
Release:    %{_rel}
Summary:    Modern docker client

Group:      containers
License:    GPL2.0
URL:        https://git.famillegratton.net:3000/devops/dtools

Source0:    %{name}-%{_version}.tar.gz
BuildArchitectures: x86_64
BuildRequires: gcc
#Requires: docker

%description
Modern-day Docker client

%prep
%autosetup

%build
cd %{_sourcedir}/%{_name}-%{_version}/src
PATH=$PATH:/opt/go/bin go build -o %{_sourcedir}/%{_binaryname} .
strip %{_sourcedir}/%{_binaryname}

%clean
rm -rf $RPM_BUILD_ROOT

%pre
if getent group devops > /dev/null; then
  exit 0
else
  if getent group 2500; then
    groupadd devops
  else
    groupadd -g 2500 devops
  fi
fi
exit 0

%install
install -Dpm 0755 %{_sourcedir}/%{_binaryname} %{buildroot}%{_bindir}/%{_binaryname}

%post
cd /opt/bin
sudo chgrp -R 0:devops .
sudo chmod 775 /opt/bin/dtools

%preun

%postun

%files
%defattr(-,root,root,-)
%{_bindir}/%{_binaryname}


%changelog
* Sat Jan 27 2024 RPM Builder <builder@famillegratton.net> 00.73.00-0
- New output for dtools lsi (jean-francois@famillegratton.net)

* Sat Jan 27 2024 RPM Builder <builder@famillegratton.net> 00.72.00-0
- Prettified dtools ls output, version bump (jean-francois@famillegratton.net)
- buildDeps + go version bumps (jean-francois@famillegratton.net)
- Renamed a file, no biggie (jean-francois@famillegratton.net)
- Fixed APK release number (builder@famillegratton.net)

* Sat Dec 23 2023 RPM Builder <builder@famillegratton.net> 00.70.00-0
- Fixed unresolved reference (jean-francois@famillegratton.net)
- Completed vol subcommands (jean-francois@famillegratton.net)
- Doc update, version bump (jean-francois@famillegratton.net)
- Completed volume rm (jean-francois@famillegratton.net)
- Completed volume ls (jean-francois@famillegratton.net)
- Completed volume ls (jean-francois@famillegratton.net)
- Extra commands stubs (jean-francois@famillegratton.net)
- Version bump (jean-francois@famillegratton.net)
- dtools volume stub (jean-francois@famillegratton.net)
- Moved CheckAPIversion() into system (jean-francois@famillegratton.net)

* Fri Dec 22 2023 RPM Builder <builder@famillegratton.net> 00.60.00-1
- Removed un-needed dependency that might break dnf (jean-
  francois@famillegratton.net)

* Fri Dec 22 2023 RPM Builder <builder@famillegratton.net> 00.60.00-0
- New release : 00.60.00 (jean-francois@famillegratton.net)
- Removed unused function parameter in API version check (jean-
  francois@famillegratton.net)
- doc update (jean-francois@famillegratton.net)
- Doc update, fixed stdin being echoed to stdout (jean-
  francois@famillegratton.net)
- Markdown shakeup by zigarn (Alexandre Garnier) (jean-
  francois@famillegratton.net)
- doc reformatting (jean-francois@famillegratton.net)
- Reverted to kludged dtools exec (jean-francois@famillegratton.net)

* Thu Dec 14 2023 RPM Builder <builder@famillegratton.net> 00.50.00-0
- New release with a dtools exec mostly working (jean-
  francois@famillegratton.net)
- Completed a somehow broken dtools exec (jean-francois@famillegratton.net)
- New command : dtools diff (jean-francois@famillegratton.net)
- Sync zenika->home (jean-francois@famillegratton.net)
- More .md updates (jean-francois@famillegratton.net)
- Doc update + output clarification on dtools lsi (jean-
  francois@famillegratton.net)
- Doc update, moved dtools info to dtools system info (jean-
  francois@famillegratton.net)
- Moved repo.ReadDefaultFile() to a type receiver (jean-
  francois@famillegratton.net)
- More workspace cleanup (jean-francois@famillegratton.net)
- Final doc fix (API version doc) (jean-francois@famillegratton.net)
- Final doc update... for now (jean-francois@famillegratton.net)
- Extra update (jean-francois@famillegratton.net)
- Doc update (jean-francois@famillegratton.net)
- Extra documentation (jean-francois@famillegratton.net)

* Thu Dec 07 2023 RPM Builder <builder@famillegratton.net> 00.31.00-0
- GO version bump, completed dockergettags and dockergetcatalog (jean-
  francois@famillegratton.net)
- Release bump (jean-francois@famillegratton.net)
- Completed network connect/disconnect; added docker API check before using the
  soft (jean-francois@famillegratton.net)

* Sun Dec 03 2023 RPM Builder <builder@famillegratton.net> 00.30.00-0
- Packaging version bump (jean-francois@famillegratton.net)
- Mostly English corrections, really (jean-francois@famillegratton.net)
- Completed network removal command (jean-francois@famillegratton.net)
- Fixed parameters parsing for the inspect subcommand (jean-
  francois@famillegratton.net)
- Minor file refactoring (jean-francois@famillegratton.net)
- Completed output of network ls (jean-francois@famillegratton.net)
- Sync Bergen-> (jean-francois@famillegratton.net)
- Completed network ls (jean-francois@famillegratton.net)
- Out of band fix for an overlooked issue (jean-francois@famillegratton.net)
- Enabled rmi, completed tag (jean-francois@famillegratton.net)

* Mon Nov 27 2023 RPM Builder <builder@famillegratton.net> 00.20.00-0
- Release number bump (jean-francois@famillegratton.net)
- Fixed issue of blindly pushing non-existing images (jean-
  francois@famillegratton.net)
- Sync bergen-> (jean-francois@famillegratton.net)
- Removed useless return params (jean-francois@famillegratton.net)
- Fixes to pull, completed broken push (jean-francois@famillegratton.net)
- Completed dtools info (jean-francois@famillegratton.net)
- Completed repo subcommands (jean-francois@famillegratton.net)
- Completed the system info subcommand (jean-francois@famillegratton.net)
- Fixed dumbass GoLand's excessive refactoring (jean-
  francois@famillegratton.net)
- ALPINE - Makefile fix (builder@famillegratton.net)
- Fixed control file (builder@famillegratton.net)


* Sat Nov 18 2023 RPM Builder <builder@famillegratton.net> 00.10.00-0
- new package built with tito

* Sat Nov 18 2023 RPM Builder <builder@famillegratton.net> 00.10.00-0
- new package built with tito

