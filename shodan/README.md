
<h1>Description:</h1><br>
Search multiple organization with Shodan.<br>
Prepare txt file with organizations names and pass it to script as argument.<br>

<b>Edit line #140 and change your api key.</b>

Needed libraries:<br>
```
gopkg.in/ns3777k/go-shodan.v1/shodan (Shodan API)
github.com/PuerkitoBio/goquery (gathering organization from Bugcrowd) (optional)
```
<h2>Example</h2>
Hosts.txt includes:<br>

  Sony<br>
  Facebook<br>
  Dropbox<br>

Run ./shodan hosts.txt

<h2>Output</h2><br>

As output script makes directory with organization's name and writes response as txt file<br>
```
/Sony
---xxx.xxx.xxx.xxx
---xxx.xxx.xxx.xxx
/Facebook
---xxx.xxx.xxx.xxx
---xxx.xxx.xxx.xxx
```
