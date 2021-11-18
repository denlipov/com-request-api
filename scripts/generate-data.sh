#!/bin/bash

for ((sidx = 0; sidx < 10; sidx++)); do
        for ((uidx = 0; uidx < 100; uidx++)); do
                req_id=$(grpc_cli call localhost:8082 CreateRequestV1 \
                        "request: { service: \"dum${sidx}\", user: \"test${uidx}\", text: \"hooo${sidx}${uidx}\" }" | grep request_id | egrep -o "[0-9]+")

#                sleep 0.3

                grpc_cli call localhost:8082 UpdateRequestV1 \
                        "request_id: $req_id, body: { text: \"some new text ${sidx} ${uidx}\" }"

#                sleep 0.3

                if [[ $[RANDOM%3] -eq 0 ]]; then
                        grpc_cli call localhost:8082 RemoveRequestV1 "request_id: $req_id"
                fi
        done
done

#grpc_cli call localhost:8082 DescribeRequestV1 'request_id: 101'
#grpc_cli call localhost:8082 ListRequestV1 ''
