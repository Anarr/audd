# audd

Golang client for audd.io (https://audd.io/)

## Getting Started
audd.io is Audio Recognition API service. The servic allows you to recognize audio files, recordings and audio streams.
This packkage helps you to integrate audd.io service in your Golang project easily.

### Prerequisites

audd pakcage written in Go 1.14 however supports Go 1.13+ (Starting in Go 1.13, module mode will be the default for all development)

```
Give examples
```

### Installing

```
go get -u github.com/Anarr/audd Run this command inside your project
```

## Running the tests

go test .

### Break down into end to end tests

Explain what these tests test and why

```
Give an example
```

### And coding style tests
To start implement package use the example below:

```
var audioClient audd.AuidoClient
audioClient.SetToken("YOUR_API_TOKEN").SetPlatform(audd.PlatformAppleMusic)
res, err := ac.Parse("AUDIO_FILE_URL")

if err != nil {
    log.Println(err)
    return
}

fmt.Println(res)
```

## Deployment

Add additional notes about how to deploy this on a live system

## Built With

* [Dropwizard](http://www.dropwizard.io/1.0.2/docs/) - The web framework used
* [Maven](https://maven.apache.org/) - Dependency Management
* [ROME](https://rometools.github.io/rome/) - Used to generate RSS Feeds

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags). 

## Authors

* **Anar Rzayev** - *Initial work* - [Anarr](https://github.com/Anarr)

See also the list of [contributors](https://github.com/Anarr/audd/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Hat tip to anyone whose code was used
* Inspiration
* etc

