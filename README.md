This is a golang based API for interacting with Mimecast.

To use, initalize some application API values:

	mimecast.SetMimeCastConfig(mimecast.MimeCastConfig{
		ApplicationId:  os.Getenv("MIMECAST_APPLICATION_ID"),
		ApplicationKey: os.Getenv("MIMECAST_APPLICATION_KEY"),
		AccessKey:      os.Getenv("MIMECAST_ACCESS_KEY"),
		SecretKey:      os.Getenv("MIMECAST_SECRET_KEY"),
	})


Then call one of the various public methods. More docs forthcoming.


