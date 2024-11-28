
0. 启动gaeger
    cd modify-ipfs-with-cluster/testdata/gaeger
    docker-compose up -d

1. init ipfs
    export IPFS_PATH=modify-ipfs-with-cluster/testdata/ipfs
    ipfs init
    ipfs config Routing.Type dht
    ipfs config Addresses.API /ip4/0.0.0.0/tcp/5001
    ipfs dag import webui.car
    ipfs config --json API.HTTPHeaders.Access-Control-Allow-Origin '["http://192.168.58.110:5001", "http://localhost:3000", "http://127.0.0.1:5001", "https://webui.ipfs.io"]'
    ipfs config --json API.HTTPHeaders.Access-Control-Allow-Methods '["PUT", "POST"]'

2. init ipfs-cluster
    export IPFS_CLUSTER_PATH=modify-ipfs-with-cluster/testdata/ipfs-cluster
    ipfs-cluster-service init
    将service.json中的monitor进行调整

ipfs 启用 tracing 步骤
1. export OTEL_EXPORTER_FILE_PATH=modify-ipfs-with-cluster/testdata/gaeger/ipfs-traces.json
2. 启动 OTEL_EXPORTER_OTLP_INSECURE=true OTEL_TRACES_EXPORTER=otlp ipfs daemon



ipfs-cluster 启用 tracing 步骤

1. ipfs-cluster-service daemon



