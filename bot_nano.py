#!/usr/bin/env python3
import requests
import socket
import json

headers = {
  'Content-Type': 'application/json'
}

apikey = "put-your-coinmarketcap-apikey"
url_nano = "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest?id=1567&convert=EUR&CMC_PRO_API_KEY={apikey}".format(apikey=apikey)

response = requests.request("GET", url_nano, headers=headers)

data = response.json()['data']['1567']['quote']['EUR']

price = data['price']

nanos = 249

botApiKey = "put-you-bot-apikey"
amount = nanos * price

r = requests.post('https://api.telegram.org/{botApiKey}/sendMessage'.format(botApiKey=botApiKey),
              data={'chat_id': '3281783', 'text': 'Current quote: {price}, current amount: {amount}'.format(price=round(price,5), amount=round(amount,2))})
