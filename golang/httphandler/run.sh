#! /bin/bash

URL="http://localhost:55057"

echo "===> GET ${URL}/fruits"
curl -i -X GET ${URL}/fruits
echo ""

echo "===> GET ${URL}/fruits/apple"
curl -i -X GET ${URL}/fruits/apple
echo ""

echo "===> POST ${URL}/fruits/apple"
curl -i -X POST -H "Content-Type: application/json" -d '{"weight": "100"}' ${URL}/fruits/apple
echo ""

echo "===> PUT ${URL}/fruits/apple"
curl -i -X PUT -H "Content-Type: application/json" -d '{"weight": "100"}' ${URL}/fruits/apple
echo ""

echo "===> GET ${URL}/fruits/apple"
curl -i -X GET ${URL}/fruits/apple
echo ""

echo "===> DELETE ${URL}/fruits/apple"
curl -i -X DELETE ${URL}/fruits/apple
echo ""

echo "===> GET ${URL}/fruits"
curl -i -X GET ${URL}/fruits
echo ""
