package pushpull

type Pusher interface {
	/*
		Push, given a path and some data, pushes that data to somewhere, using the string as a location and name identifier
		i.e.: Push("/test/me", `DapperDodo`) could write a file called 'me' in the folder '/test' on some ftp server with the contents 'DapperDodo'
	*/
	Push(string, []byte) error
}

type Puller interface {
	/*
		Pull, given a path, pulls data from somewhere, using the string as a location and name identifier
		i.e.: Pull("/test/me") could read a file called 'me' in the folder '/test' on some ftp server, returning its contents as an array of bytes `DapperDodo`
	*/
	Pull(string) ([]byte, error)
}

type PushPuller interface {
	Pusher
	Puller
}
