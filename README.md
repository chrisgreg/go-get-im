# go-get-im 📸 🌅

### What?

A command-line Go script to download Reddit user submitted imgur albums and files using the Reddit & Imgur APIs.

Results will be added to a local folder `go-get-im-results` namespaced by username.

### Reguirements
- The Reddit Username you want to download imgur submissions of

### Installation
`go build go-get-im.go`

### Usage
`./go-get-im "reddit-user-name-goes-here"`


## Todo
- Wrap some calls in Goroutines for performance
- Tests
