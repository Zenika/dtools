- [x] stopall, startall, etc will try to stop/start containers even when already stopped or start<br> 
- [x] dtools pull prepends a / before the args (??)
- [ ] more work on `repo add` to prompt for values
- [x] better error handling (connection refused on `pull`)
- [ ] Multiple prints of errors with `pull` (<-- needs re-investigation, unsure it is a valid use)
- [ ] "prettify" `repo ls` <<-- it's good enough for now.
- [x] complete the so-far sparse `system info` subcommand
- [ ] push is not pushing anything
- [x] push reports success pushing an image that does not even exist ( !! )
- [ ] minimal Docker API version is currently hardcoded in main() 
- [x] Refactor all functions in the repo package, so they use type receivers
- [ ] `dtools run`
- [ ] `dtools network add`
- [ ] "intelligent" docker system prune
- [x] META: prettify MAPPINGS.md
- [ ] `dtools exec` : issues with stdin sent to stdout, CTRL+D not exiting shells
- [ ] corner case where `dtools rmi` removes an image when a container is running ? <-- needs investigation

<br><br><br>
