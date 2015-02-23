a2csv
====

## Description
Output the information of AWS in CSV format.

At the moment, the supported resources are as below.

  - SecurityGroup ( the Detailed rules of inbound and outbound )

## Usage

```
./a2csv -h
./a2csv securityGroups -r [ region ] -p [ profile ]
```

```
$ ./a2csv securityGroups -r ap-northeast-1 -p private
GroupID,Direction,Type,Protocol,Port Range,IP Range
sg-b6822ed3,inbound,ALL ICMP,ICMP(1),ALL,sg-a9822ecc
sg-b6822ed3,outbound,ALL Traffic,ALL,ALL,0.0.0.0/0
sg-6ad86e0f,inbound,SSH(22),TCP(6),22,0.0.0.0/0
sg-6ad86e0f,outbound,ALL Traffic,ALL,ALL,0.0.0.0/0
sg-f07cdd95,inbound,RDP(3389),TCP(6),3389,192.168.11.1/32
sg-f07cdd95,outbound,ALL Traffic,ALL,ALL,0.0.0.0/0
sg-a9822ecc,inbound,ALL Traffic,ALL,ALL,192.168.11.1/32
sg-a9822ecc,outbound,ALL ICMP,ICMP(1),ALL,sg-b6822ed3
sg-aa2686cf,inbound,Custom Rrotocol Rule,HOPOPT(0),ALL,sg-6ad86e0f
sg-aa2686cf,inbound,SSH(22),TCP(6),22,1.1.1.1/32
sg-aa2686cf,outbound,ALL Traffic,ALL,ALL,0.0.0.0/0
sg-b10aecd4,inbound,SSH(22),TCP(6),22,0.0.0.0/0
sg-b10aecd4,outbound,ALL Traffic,ALL,ALL,0.0.0.0/0
```

## Install
Git clone and go build. 
To suit your environment, please change the options of `go build`

```bash
$ git clone https://github.com/kongou-ae/aws-to-csv.git
$ go build a2csv.go commands.go version.go security-groups-detail.go
$ vi ~/.aws/credentials
```

And input your information of AWS as below.

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
