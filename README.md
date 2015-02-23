a2csv
====

## Description
output the information of AWS in CSV format.
At the moment, resources that are supported are as below.

  - SecurityGroup(the Detailed rules of inbound and outbound)

## Usage

./a2csv securityGroups -r [region] -p [profile]

## Install
git clone and go build. 
To suit your environment, please change the options of "go build"

```bash
$ git clone https://github.com/kongou-ae/aws-to-csv.git
$ go build a2csv.go commands.go version.go security-groups-detail.go
$ vi ~/.aws/credentials
```

and input your information of AWS as below.

```
[PROFILE]
aws_access_key_id = YOUR_ACCESS_KEY
aws_secret_access_key = YOUR_SECRET_KEY
```

## Contribution

1. Fork ([https://github.com/kongou-ae/a2csv/fork](https://github.com/kongou-ae/a2csv/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create new Pull Request

## Author

[kongou-ae](https://github.com/kongou-ae)
