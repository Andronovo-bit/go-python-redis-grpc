# GO-PYTHON-GRPC-REDIS

## Giriş

Bu projede docker sanallaştırması ve gRPC protokolü kullanılarak json içerisinde bulunan veriler go dili ile yazılmış client kod bloğu ile okunarak protobuf'a çevirlierek pyhton dili ile yazılmış server'e gönderilir. Daha sonra server ise gelen verileri alarak redis'e kayıt eder.

Projenin çalışabilmesi için ilk öncelikle bilgisayarınızda docker kurulu olmalıdır.

Windows için --> https://download.docker.com/win/stable/Docker%20Desktop%20Installer.exe

Docker kurulumu yapıldıktan sonra projeyi bilgisayarınıza çekebilirsiniz.

```git clone https://github.com/Andronovo-bit/go-python-redis-grpc.git```

## Kurulum

İlk olarak proje ana dizinine girin.

1-) ```docker-compose config``` komutu ile config dosyasını kontrol ediniz

2-) ```docker-compose up --build -d```  komutu ile server ve redis containerlerini oluşturunuz.

3-) ` cd client ` komutu ile ana dizinden client dizinine geçiş yapınız.

4-)  `docker build -t client . -f dockerfile  ` komutu ile client projesini build ediniz.


## Çalıştırma

Projeyi çalıştımak için

1-) ```docker run -it client``` komutu ile client projesini çalıştırarak containeri çalışır hale getiriniz.

2-)  Aynı terminal ekranında  ``` cd client ``` komutu ile proje dosyasının olduğu dizine geçiniz.

3-)  ```go run main.go``` komutu ile projeyi çalıştırabilirsiniz.

Proje çalıştıktan sonra terminal ekranında json'daki verilerin kayıt edildiğini göreceksiniz. Redis'e giderek kayıt edilen verileri görebilirsiniz. Key olarak ip adresi kullanılmıştır. 

---------------------------ENGLISH------------------------------------

## Introduction

In this project, using the docker virtualization and gRPC protocol, the data in json is read with the client code block written in go language and translated into protobuf and sent to the server written in pyhton language. Then the server takes the incoming data and saves it to redis.

In order for the project to work, you must first have docker installed on your computer.

For Windows --> https://download.docker.com/win/stable/Docker%20Desktop%20Installer.exe

Once the Docker is installed, pull the project onto your computer.

```git clone https://github.com/Andronovo-bit/go-python-redis-grpc.git```

## Installation

First, enter the project home directory.

1-) Check the config file with the command ``` docker-compose config ```.

2-) Create the server and redis containers with the command ``` docker-compose up --build -d ```.

3-) Switch from the main directory to the client directory with the  ```cd client ``` command.

4-) Build the client project with the command ```docker build -t client . -f dockerfile```.


## Run

To run the project;

1-) ```docker run -it client``` komutu ile client projesini çalıştırarak containeri çalışır hale getiriniz.

2-) Make the container run by running the client project with the command ```docker run -it client```.

3-) You can run the project with the command ``` go run main.go ```.

After the project is running, you will see that the data in json is saved on the terminal screen. You can see the recorded data by going to Redis. IP address is used as key.
