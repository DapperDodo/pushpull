# pushpull

`push/pull data to/from any(tm) server/filesystem`

Pushpull is a convenience adapter for writing and reading data (files) to and from polymorph sources.


#### rationale:

In my daily job I am frequently asked to import or export `stuff` to some location shared with the customer.
These locations vary a lot in type. ftp, sftp, ftps, webdav etc...
But the tasks are almost always the same: 
- Export some generated file to the customer. 
- Import data from some file on the customer's server.

It occurred to me that >80% of such cases could be solved in <20% of custom development time if I had a tool that would:
- Abstract away the type of remote location
- Provide simple configuration for any chosen backend type
- Just Work(tm)!!


#### interface:

pushpull is designed around two very simple interfaces:

	push(filepath, []byte) error
	pull(filepath) ([]byte, error)


#### installation:

	go get -u github.com/DapperDodo/pushpull


#### backends:

	ftp
	ftps (coming soon)
	...

#### wishlist:

	sftp
	scp
	webdav
	amazon s3
	google cloud storage
	dropbox
	...


#### basic example main.go:

	import (
		"fmt"
		
		"github.com/DapperDodo/pushpull"
	)
	
	func main() {

		// our server
		var server pushpull.PushPuller

		// now point it to a backend instance, in this case a simple ftp server
		srv := "localhost"
		prt := "21"
		usr := "username"
		pwd := "password"
		server := pushpull.NewFtp(srv, prt, usr, pwd)

		// imagine you just generated some valuable data
		data_generated := `hello world!`

		// push that data to your customer's server like this:
		_ = server.Push('/some/path/filename', data_generated)

		// to prove it works, pull that same data from the server like this:
		data_pull, _ := server.Pull('/some/path/filename')

		// in real applications, handle errors appropriately!

		fmt.Println("to yonder and back!:", string(data_pull))
	}
