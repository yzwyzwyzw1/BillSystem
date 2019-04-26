#/bin/bahs!
export CHANNEL_NAME=mychannel
echo "创建通道"
peer channel create -o orderer.example.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/mychannel.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
echo "节点加入通道"
peer channel join -b mychannel.block
echo "安装链码"
peer chaincode install -n mycc -v 1.0 -p github.com/chaincode
echo "实例化链码"
peer chaincode instantiate -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n mycc -v 1.0 -c'{"Args":["init"]}' -P "OR ('org1.example.com.member')"
echo "睡眠5s,等待实例化成功"
sleep 5s
echo "发布票据BOC1001"
peer chaincode invoke -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n mycc  -c '{"Args":["issue","BOC1001","20000","111","20180101","20180111","111","111","111","111","111","111","jack","jackID"]}'
sleep 5s
echo "发布票据BOC1002"
peer chaincode invoke -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n mycc  -c '{"Args":["issue","BOC1002","10000","111","20180101","20180111","111","111","111","111","111","111","jack","jackID"]}'
sleep 5s
echo "查询jack的票据"
peer chaincode query -C $CHANNEL_NAME -n mycc -c '{"Args":["queryBills","jackID"]}'

sleep 5s
echo "jack向alice发起签收票据请求"
peer chaincode invoke -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n mycc  -c '{"Args":["endorse","BOC1001","alice","aliceID"]}'
sleep 5s
echo "alice签收票据"
peer chaincode invoke -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n mycc  -c '{"Args":["accept","BOC1001","alice","aliceID"]}'
sleep 5s
echo "查询票据BOC1001"
peer chaincode query -C $CHANNEL_NAME -n mycc -c '{"Args":["queryBillByNo","BOC1001"]}'

sleep 5s
echo "jack向alice发起签收票据请求"
peer chaincode invoke -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n mycc  -c '{"Args":["endorse","BOC1002","alice","aliceID"]}'
sleep 5s
echo "alice拒绝签收票据"
peer chaincode invoke -o orderer.example.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n mycc  -c '{"Args":["reject","BOC1002","alice","aliceID"]}'
sleep 5s
echo "查询jack的票据"
peer chaincode query -C $CHANNEL_NAME -n mycc -c '{"Args":["queryBills","jackID"]}'
sleep 5s
echo "查询票据BOC1002"
peer chaincode query -C $CHANNEL_NAME -n mycc -c '{"Args":["queryBillByNo","BOC1002"]}'
