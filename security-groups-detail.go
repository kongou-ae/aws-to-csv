package main

import (
  "fmt"
  "time"
  "net/http"
  "strconv"
  "github.com/awslabs/aws-sdk-go/aws"
  "github.com/awslabs/aws-sdk-go/gen/ec2"
)

// メイン処理
func SecurityGroupsDetail(prefix string,region string) {

  // ~/.aws/credentials からキーを取得
  duration := 60 * time.Second
  creds, err := aws.ProfileCreds("",prefix,duration)
  if err != nil {
    fmt.Print("Error!! Please check ~/.aws/credentials")
  }

  // 取得したキーとコマンドラインからの入力を利用して、APIへ接続
  cli := ec2.New(creds, region, http.DefaultClient)
  // aws ec2 describe-security-grroupsを実行
  resp, err := cli.DescribeSecurityGroups(nil)

  if err != nil {
    panic(err)
  }

  // セキュリティグループの出力処理
  fmt.Println("GroupID,Direction,Type,Protocol,Port Range,IP Range")
  // ループを利用して、セキュリティグループ一つずつに処理を実行する
  for i := range resp.SecurityGroups {
    // ループを利用して、1つのセキュリティグループ内のinboundポリシーに処理を実行する
    if resp.SecurityGroups[i].IPPermissions != nil {
      for j := range resp.SecurityGroups[i].IPPermissions {
        // ターゲットがIPレンジとセキュリティグループのとちらかチェック。
        if resp.SecurityGroups[i].IPPermissions[j].IPRanges != nil {
          // 送信元IPアドレス分処理を実施
          for k := range resp.SecurityGroups[i].IPPermissions[j].IPRanges {
            fmt.Print(*resp.SecurityGroups[i].GroupID + ",inbound,")
            // ルールの詳細を出力
            print_detail(resp.SecurityGroups[i].IPPermissions[j])
            fmt.Print(*resp.SecurityGroups[i].IPPermissions[j].IPRanges[k].CIDRIP + "\n")
          }
        } else {
          // 送信元セキュリティグループ分処理を実施
          for k := range resp.SecurityGroups[i].IPPermissions[j].UserIDGroupPairs {
            fmt.Print(*resp.SecurityGroups[i].GroupID + ",inbound,")
            // ルールの詳細を出力
            print_detail(resp.SecurityGroups[i].IPPermissions[j])
            fmt.Print(*resp.SecurityGroups[i].IPPermissions[j].UserIDGroupPairs[k].GroupID + "\n")
          }
        }
      }
    }
    // ループを利用して、1つのセキュリティグループ内のoutboundポリシーに処理を実行する
    if resp.SecurityGroups[i].IPPermissionsEgress != nil {
      for j := range resp.SecurityGroups[i].IPPermissionsEgress {
        // ターゲットがIPレンジとセキュリティグループのとちらかチェック。
        if resp.SecurityGroups[i].IPPermissionsEgress[j].IPRanges != nil {
          // 送信元IPアドレス分処理を実施
          for k := range resp.SecurityGroups[i].IPPermissionsEgress[j].IPRanges {
            fmt.Print(*resp.SecurityGroups[i].GroupID + ",outbound,")
            // ルールの詳細を出力
            print_detail(resp.SecurityGroups[i].IPPermissionsEgress[j])
            fmt.Print(*resp.SecurityGroups[i].IPPermissionsEgress[j].IPRanges[k].CIDRIP + "\n")
          }
        } else {
          // 送信元セキュリティグループ分処理を実施
          for k := range resp.SecurityGroups[i].IPPermissionsEgress[j].UserIDGroupPairs {
            fmt.Print(*resp.SecurityGroups[i].GroupID + ",outbound,")
            // ルールの詳細を出力
            print_detail(resp.SecurityGroups[i].IPPermissionsEgress[j])
            fmt.Print(*resp.SecurityGroups[i].IPPermissionsEgress[j].UserIDGroupPairs[k].GroupID + "\n")
          }
        }
      }
    }
  }
}

func print_detail (sg_rule ec2.IPPermission) {

  var (
   fromPort int = 0
   toPort int = 0
   portRange string = ""
  )


  icmp_code := map[string]string{
      "-1":"ALL",
      "3--1":"Destination Unreachable / All",
      "3-0":"Destination Unreachable / destination network unreachable",
      "3-1":"Destination Unreachable / destination host unreachable",
      "3-2":"Destination Unreachable / destination protocol unreachable",
      "3-3":"Destination Unreachable / destination port unreachable",
      "3-4":"Destination Unreachable / fragmentation required, and?DF flag?set",
      "3-5":"Destination Unreachable / source route failed",
      "3-6":"Destination Unreachable / destination network unknown",
      "3-7":"Destination Unreachable / destination host unknown",
      "3-8":"Destination Unreachable / source host isolated",
      "3-9":"Destination Unreachable / network administratively prohibited",
      "3-10":"Destination Unreachable / host administratively prohibited",
      "3-11":"Destination Unreachable / network unreachable for TOS",
      "3-12":"Destination Unreachable / host unreachable for TOS",
      "3-13":"Destination Unreachable / communication administratively prohibited",
      "4--1":"Source quench (congestion control)",
      "5--1":"Redirect Message / All",
      "5-0":"Redirect Message / redirect datagram for the network",
      "5-1":"Redirect Message / redirect datagram for the host",
      "5-2":"Redirect Message / redirect datagram for the TOS & network",
      "5-3":"Redirect Message / redirect datagram for the TOS & host",
      "6--1":"Alternate Host Address",
      "7--1":"Reserved",
      "8--1":"Echo request (used to ping)",
      "9--1":"Router Advertisement",
      "10--1":"Router Solicitation",
      "11--1":"Time Exceeded / All",
      "11-0":"Time Exceeded / TTL expired in transit",
      "11-1":"Time Exceeded / fragment reassembly time exceeded",
      "12--1":"Parameter Problem: Bad IP Header / All",
      "12-0":"Parameter Problem: Bad IP Header / pointer indicates the error",
      "12-1":"Parameter Problem: Bad IP Header / missing a required option",
      "12-2":"Parameter Problem: Bad IP Header / bad length",
      "13--1":"Timestamp",
      "14--1":"Timestamp reply",
      "15--1":"Information Request",
      "16--1":"Information Reply",
      "17--1":"Address Mask Request",
      "18--1":"Address Mask Reply",
      "30--1":"Traceroute",
      "31--1":"Datagram Conversion Error",
      "32--1":"Mobile Host Redirect",
      "33--1":"Where Are You",
      "34--1":"Here I Am",
      "35--1":"Mobile Registration Request",
      "36--1":"Mobile Registration Reply",
      "37--1":"Domain Name Request",
      "38--1":"Domain Name Reply",
      "39--1":"SKIP Algorithm Discovery Protocol",
      "40--1":"Photuris, Security failures",
  }

  protocol := map[string]string{
      "0":"HOPOPT",
      "1":"ICMP",
      "2":"IGMP",
      "3":"GGP",
      "4":"IP",
      "5":"ST",
      "6":"TCP",
      "7":"CBT",
      "8":"EGP",
      "9":"IGP",
      "10":"BBN-RCC-MON",
      "11":"NVP-II",
      "12":"PUP",
      "13":"ARGUS",
      "14":"EMCON",
      "15":"XNET",
      "16":"CHAOS",
      "17":"UDP",
      "18":"MUX",
      "19":"DCN-MEAS",
      "20":"HMP",
      "21":"PRM",
      "22":"XNS-IDP",
      "23":"TRUNK-1",
      "24":"TRUNK-2",
      "25":"LEAF-1",
      "26":"LEAF-2",
      "27":"RDP",
      "28":"IRTP",
      "29":"ISO-TP4",
      "30":"NETBLT",
      "31":"MFE-NSP",
      "32":"MERIT-INP",
      "33":"DCCP",
      "34":"3PC",
      "35":"IDPR",
      "36":"XTP",
      "37":"DDP",
      "38":"IDPR-CMTP",
      "39":"TP++",
      "40":"IL",
      "41":"IPv6",
      "42":"SDRP",
      "43":"IPv6-Route",
      "44":"IPv6-Frag",
      "45":"IDRP",
      "46":"RSVP",
      "47":"GRE",
      "48":"MHRP",
      "49":"BNA",
      "50":"ESP",
      "51":"AH",
      "52":"I-NLSP",
      "53":"SWIPE",
      "54":"NARP",
      "55":"MOBILE",
      "56":"TLSP",
      "57":"SKIP",
      "58":"IPv6-ICMP",
      "59":"IPv6-NoNxt",
      "60":"IPv6-Opts",
      "61":"-",
      "62":"CFTP",
      "63":"-",
      "64":"SAT-EXPAK",
      "65":"KRYPTOLAN",
      "66":"RVD",
      "67":"IPPC",
      "68":"-",
      "69":"SAT-MON",
      "70":"VISA",
      "71":"IPCV",
      "72":"CPNX",
      "73":"CPHB",
      "74":"WSN",
      "75":"PVP",
      "76":"BR-SAT-MON",
      "77":"SUN-ND",
      "78":"WB-MON",
      "79":"WB-EXPAK",
      "80":"ISO-IP",
      "81":"VMTP",
      "82":"SECURE-VMTP",
      "83":"VINES",
      "84":"TTP",
      "85":"NSFNET-IGP",
      "86":"DGP",
      "87":"TCF",
      "88":"EIGRP",
      "89":"OSPFIGP",
      "90":"Sprite-RPC",
      "91":"LARP",
      "92":"MTP",
      "93":"AX.25",
      "94":"IPIP",
      "95":"MICP",
      "96":"SCC-SP",
      "97":"ETHERIP",
      "98":"ENCAP",
      "99":"-",
      "100":"GMTP",
      "101":"IFMP",
      "102":"PNNI",
      "103":"PIM",
      "104":"ARIS",
      "105":"SCPS",
      "106":"QNX",
      "107":"A/N",
      "108":"IPComp",
      "109":"SNP",
      "110":"Compaq-Peer",
      "111":"IPX-in-IP",
      "112":"VRRP",
      "113":"PGM",
      "114":"-",
      "115":"L2TP",
      "116":"DDX",
      "117":"IATP",
      "118":"STP",
      "119":"SRP",
      "120":"UTI",
      "121":"SMP",
      "122":"SM",
      "123":"PTP",
      "124":"ISIS over IPv4",
      "125":"FIRE",
      "126":"CRTP",
      "127":"CRUDP",
      "128":"SSCOPMCE",
      "129":"IPLT",
      "130":"SPS",
      "131":"PIPE",
      "132":"SCTP",
      "133":"FC",
      "134":"RSVP-E2E-IGNORE",
      "135":"Mobility Header",
      "136":"UDPLite",
      "137":"MPLS-in-IP",
      "138":"manet",
      "139":"HIP",
      "140":"Shim6",
      "141":"WESP",
      "142":"ROHC",
      "253":"-",
      "254":"-",
  }

  sg_type := map[string]string {
      "tcp":"Custom TCP Rule",
      "udp":"Custom UDP Rule",
      "icmp":"Custom ICMP Rule",
      "custom":"Custom Rrotocol Rule",
      "tcp/ALL":"ALL TCP",
      "udp/ALL":"ALL UDP",
      "icmp/-1":"ALL ICMP",
      "-1":"ALL Traffic",
      "tcp/22":"SSH(22)",
      "tcp/23":"telnet(23)",
      "tcp/25":"SMTP(25)",
      "tcp/42":"nameserver(42)",
      "udp/53":"DNS(53)",
      "tcp/80":"HTTP(80)",
      "tcp/110":"POP3(110)",
      "tcp/143":"IMAP(143)",
      "tcp/389":"LDAP(389)",
      "tcp/443":"HTTPS(443)",
      "tcp/465":"SMTP(465)",
      "tcp/993":"IMAPS(993)",
      "tcp/9955":"POP3S(995)",
      "tcp/1433":"MS SQL(1433)",
      "tcp/3306":"MySQL(3306)",
      "tcp/3389":"RDP(3389)",
      "tcp/8080":"HTTP*(8086)",
      "tcp/8443":"HTTPS*(8443)",
   }

   // APIの結果にFromPortがあるかチェック
   if sg_rule.FromPort != nil {
     fromPort = *sg_rule.FromPort
   }
   // APIの結果にToPortがあるかチェック
   if sg_rule.ToPort != nil {
     toPort = *sg_rule.ToPort
   }

   // FromとToが同じであれば、出力にFromを利用する。もし違うなら、From-Toの形式を利用する
   if fromPort == toPort {
     portRange = strconv.Itoa(fromPort)
   } else {
     portRange = strconv.Itoa(fromPort) + "-" + strconv.Itoa(toPort)
   }

   // PortRangeが0-65535なら、出力をALLに変更。IFでいいかも。
   switch portRange {
   case "0-65535":
     portRange = "ALL"
   }

  // プロトコルごとに処理を変更
  switch *sg_rule.IPProtocol {
  // ALL Trafficの場合
  case "-1":
    fmt.Print(sg_type[*sg_rule.IPProtocol] + ",ALL,ALL,")
  // TCPの場合
  case "tcp":
    // Typeを出力
    if sg_type[*sg_rule.IPProtocol + "/" + portRange] == "" {
      fmt.Print(sg_type[*sg_rule.IPProtocol] + ",")
    } else {
      fmt.Print(sg_type[*sg_rule.IPProtocol + "/" + portRange] + ",")
    }
    // Protocolを出力
    fmt.Print("TCP(6),")
    fmt.Print(portRange +",")
  // UDPの場合
  case "udp":
    // Typeを出力$
    if sg_type[*sg_rule.IPProtocol + "/" + portRange] == "" {
      fmt.Print(sg_type[*sg_rule.IPProtocol] + ",")
    } else {
      fmt.Print(sg_type[*sg_rule.IPProtocol + "/" + portRange] + ",")
    }
    // Protocolを出力
    fmt.Print("UDP(17),")
    fmt.Print(portRange +",")
  // ICMPの場合
  case "icmp":
    // Typeを出力$
    if sg_type[*sg_rule.IPProtocol + "/" + portRange] == "" {
      fmt.Print(sg_type[*sg_rule.IPProtocol] + ",")
    } else {
      fmt.Print(sg_type[*sg_rule.IPProtocol + "/" + portRange] + ",")
    }
    // Protocolを出力
    fmt.Print("ICMP(1),")
    fmt.Print(icmp_code[portRange] +",")
  // custom protocolの場合
  default:
    // Typeを出力
    fmt.Print(sg_type["custom"] + ",")
    // Protocolを出力
    if protocol[*sg_rule.IPProtocol] == "-" {
      fmt.Print(*sg_rule.IPProtocol + ",")
    } else {
      fmt.Print(protocol[*sg_rule.IPProtocol] + "(" + *sg_rule.IPProtocol + "),")
    }
    // Protocolを出力
    fmt.Print("ALL,")
  }
}
