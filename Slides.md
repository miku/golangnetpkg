# Networking with Go

A short tour through the net and net/http packages.

* Martin Czygan - gh/miku, @cvvfj, martin.czygan@gmail.com
* Spartakiade 2019-04-28, 10:00, Berlin, Alte Börse Marzahn

# Welcome

In the next few hours, we will look and implement networking and web examples in
Go with net, net/http and a few helper packages.

# Outline

## Lookups and TCP servers (5)

* Basics: IP addresses, masks (ipaddr, ipmask)
* Lookups: Hosts, name resolution (lookuphost, lookupport)
* A HTTP HEAD request via TCP (tcphttphead, tcphttpheadi)
* Daytime: A daytime server (daytime, daytimei)
* Echo: An echoserver, single and multi-threaded (*)

## Fetching data (5)

* Fetching HTTP: basics, creating requests (furl1)
* http.Response (httpresponse)
* DumpResponse (dumpresponse)
* A sequential link checker (furl2, furl3)
* Fetching data in parallel (furl4)

## Go in a Container (1)

* Fetch FROM SCRATCH: putting Go into a Container (fromscratch)

## Web applications (13)

* HelloWorld (helloworld)
* Handlers (handlers)
* HandlerFunc (handlerfunc)
* Multiple handlers (multiplehandlers)
* Custom Server (customhandler)
* Fileserver (fileserver, fileservercont)
* http.Request (httprequest)
* Serializing Data (servejson)
* TLS (tls)
* testing (responserecorder)

## Web Framework (2)

* gorilla/mux (gorillamux)
* struct handler (groupstruct)

## Proxy example (1)

* Minimal HTTP proxy example (httpproxy)
