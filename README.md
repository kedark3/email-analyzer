# Emails-Analyzer

This repository contains a Go program that can be used to analyze emails.

For example, given the input:

    Delivered-To: santa@northpole.com

    Received: by 10.159.41.68 with SMTP id t62csp570647uat;

        Thu, 16 Mar 2017 04:44:28 -0700 (PDT)

    X-Received: by 10.99.140.69 with SMTP id q5mr9342725pgn.179.1489664668601;

        Thu, 16 Mar 2017 04:44:28 -0700 (PDT)

    Return-Path: <llbean.511020444@envfrm.r1234s5.com>

    Received: from omp.e1.llbean.com (omp.e1.llbean.com. [199.7.202.38])

        by mx.google.com with ESMTPS id d2si234254576pli.110.2017.03.16.04.44.28

        for <santa@northpole.com>

        (version=TLS1_2 cipher=ECDHE-RSA-AES128-GCM-SHA256 bits=128/128);

        Thu, 16 Mar 2017 04:44:28 -0700 (PDT)

    Received-SPF: pass (google.com: domain of llbean.511020444@envfrm.r1234s5.com designates 199.8.123.38 as permitted sender) client-ip=199.8.123.38;

    Authentication-Results: mx.google.com;

        dkim=pass header.i=@e1.llbean.com;

        dkim=pass header.i=@responsys.net;

        dkim=pass header.i=@responsys.net;

        spf=pass (google.com: domain of llbean.511020444@envfrm.r1234s5.com designates 199.8.123.38 as permitted sender) smtp.mailfrom=llbean.511020444@envfrm.r1234s5.com;

        dmarc=pass (p=REJECT sp=REJECT dis=NONE) header.from=e1.llbean.com

    Received: by omp.e1.llbean.com id hp9t9o1hctok for <santa@northpole.com>; Thu, 16 Mar 2017 04:25:51 -0700 (envelope-from <llbean.511020444@envfrm.r1234s5.com>)

    X-CSA-Complaints: whitelist-complaints@eco.de

    Received: by omp.e1.llbean.com id hp9r3u1hctou for <santa@northpole.com>; Thu, 16 Mar 2017 04:22:00 -0700 (envelope-from <llbean.50094@envfrm.rsys5.com>)

    MIME-Version: 1.0

    Content-Type: multipart/mixed; boundary="----msg_border_dPN32EhfaN"

    Date: Thu, 16 Mar 2017 04:22:00 -0700

    To: santa@northpole.com

    From: "L.L.Bean" <llbean@e1.llbean.com>

    Reply-To: "L.L.Bean" <reply@e1.llbean.com>

    Subject: Climbing mountains. Breaking barriers.

    Feedback-ID: 50021:10567834:oraclersys

    Message-ID: <0.0.4D.537.1D2EDF347E56956.0@omp.e1.llbean.com>

    

    ------msg_border_dPN32EhfaN

    Date: Thu, 16 Mar 2017 04:22:00 -0700

    Content-Type: multipart/alternative; boundary="----alt_border_BrDQaPYzhN_1"

    

    ------alt_border_BrDQaPYzhN_1

    Content-Type: text/plain; charset="UTF-8"

    Content-Transfer-Encoding: quoted-printable

    

    L.L.Bean Email Update

    ____________________________________________________________

    To view this email with images, click below.

The JSON returned will be:

    {

    "To": "santa@northpole.com",

    "From": "\"L.L.Bean\" <llbean@e1.llbean.com>",

    "Date": "Thu, 16 Mar 2017 04:22:00 -0700",

    "Subject": "Climbing mountains. Breaking barriers.",

    "Message-ID": "<0.0.4D.537.1D2EDF347E56956.0@omp.e1.llbean.com>"

    }


## To run this program, you can use following 3 ways:

### Method 1 - Run locally:

1. Clone the git repo using `go get`
```sh
go get github.com/kedark3/email-analyzer
```
2. Under your GOPATH, find the directory where the project is cloned. Usually at `/home/<username>/go/src/github.com/kedark3/email-analyzer` and `cd` into that directory
3. Use `go run main.go` and this should start the application with port `8080` exposed. You can then access the application at http://localhost:8080/api/v1/emails


### Method 2 - Run in a Container :

1. Clone the git repo using `go get`
```sh
go get github.com/kedark3/email-analyzer
```
2. Under your GOPATH, find the directory where the project is cloned. Usually at `/home/<username>/go/src/github.com/kedark3/email-analyzer` and `cd` into that directory
3. Use Docker or Podman to build the container locally. I have tested it with `podman` on `Fedora 32` as:
```sh
podman build -t email-analyzer . --cgroup-manager=cgroupfs
```
4. Once container is built, run it as follows:
```sh
podman run --rm -p 8080:8080 email-analyzer
```
5. You can then access the application at http://localhost:8080/api/v1/emails

### Method 3 - Run on Kubernetes:

In this method, I have tested the container running with Minikube. But if you have K8s/OpenShift cluster and Ingress controller configured, you should be able to use that as well with the files in the `k8s-configs`.

1. Clone the git repo using `go get`
```sh
go get github.com/kedark3/email-analyzer
```
2. Under your GOPATH, find the directory where the project is cloned. Usually at `/home/<username>/go/src/github.com/kedark3/email-analyzer` and `cd` into that directory
3. Authenticate to your K8s cluster and run:
```sh
kubectl create -f k8s-configs/
```
4. Find the ingress resource created in your namespace:
```sh
kubectl get ingress
```
5. Once Ingress resource gets an address and ready for use, head over to `http://emails-analyzer.com/api/v1/emails` and start using the api.


## How do I POST to API?


```sh

curl 'localhost:8080/api/v1/emails' --header 'Content-Type: text/plain' --data-binary @sampleEmails/20110401_1000mercis_14461469_html.msg | jq 

 % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 47593  100   350  100 47243   113k  15.0M --:--:-- --:--:-- --:--:-- 15.1M
{
  "To": "1000mercis@cp.assurance.returnpath.net",
  "From": "\"Darty\" <infos@contact-darty.com>",
  "Date": "01 Apr 2011 16:17:41 +0200",
  "Subject": "Cuit Vapeur 29.90 euros, Nintendo 3DS 239 euros, GPS TOM TOM 139 euros... decouvrez VITE tous les bons plans du weekend !",
  "Message-ID": "<20110401161739.E3786358A9D7B977@contact-darty.com>"
}

```
If you use the K8s, please use appropriate URL instead of `localhost`.

Use `data-binary` to preserve newlines.

You can also use API client like Postman.


## Testing

There is a `test.go` file under `tests/` directory, you can run that as `go run tests/test.go` and that will run the tests using all the files in the `sampleEmails/` and show you the output.