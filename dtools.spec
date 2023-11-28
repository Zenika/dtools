%define debug_package   %{nil}
%define _build_id_links none
%define _name dtools
%define _prefix /opt
%define _version 00.20.00
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
Requires: docker

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
exit 0

%install
install -Dpm 0755 %{_sourcedir}/%{_binaryname} %{buildroot}%{_bindir}/%{_binaryname}

%post

%preun

%postun

%files
%defattr(-,root,root,-)
%{_bindir}/%{_binaryname}


%changelog
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

