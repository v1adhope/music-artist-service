#!/bin/bash

# INFO: test tls
grpcurl -cacert certs/ca_cert.pem  localhost:50051 list
echo -e "\n"

# INFO: test no tls
grpcurl -insecure localhost:50051 list
echo -e "\n"

# INFO: examples
grpcurl -insecure\
  -d '{"data": {"name": "Some name", "description": "Some long desc", "website": "http://facebook.com/someprofile", "mounthlyListeners": 10000000, "email": "example@foo.com"}}'\
  localhost:50051\
  artist.Artist.Create

grpcurl -insecure\
  -d '{"data": {"id": "1ef5ad00-8df9-6f50-9f80-0187f4d47e03"}}'\
  localhost:50051\
  artist.Artist.Get

grpcurl -insecure\
  localhost:50051\
  artist.Artist.GetAll

grpcurl -insecure\
  -d '{"data": {"id": "1ef5a5ca-1307-6240-a42a-7727cd5c9131", "name": "Some second name", "description": "Some desc2", "website": "http://facebook.com/someprofile2", "mounthlyListeners": 20000000, "email": "example2@foo.com"}}'\
  localhost:50051\
  artist.Artist.Replace

grpcurl -insecure\
  -d '{"data": {"id": "1ef5a5ca-1307-6240-a42a-7727cd5c9131"}}'\
  localhost:50051\
  artist.Artist.Delete
